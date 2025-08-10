package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// FileInfo chứa thông tin về file và số dòng
type FileInfo struct {
	Path  string
	Lines int
}

// isCodeFile kiểm tra xem file có phải là file code không
func isCodeFile(filename string) bool {
	codeExtensions := map[string]bool{
		".go":    true,
		".js":    true,
		".ts":    true,
		".py":    true,
		".java":  true,
		".c":     true,
		".cpp":   true,
		".h":     true,
		".hpp":   true,
		".cs":    true,
		".php":   true,
		".rb":    true,
		".rs":    true,
		".swift": true,
		".kt":    true,
		".scala": true,
		".r":     true,
		".m":     true,
		".pl":    true,
		".sh":    true,
		".bash":  true,
		".zsh":   true,
		".fish":  true,
		".ps1":   true,
		".html":  true,
		".htm":   true,
		".css":   true,
		".scss":  true,
		".sass":  true,
		".less":  true,
		".vue":   true,
		".jsx":   true,
		".tsx":   true,
		".xml":   true,
		".json":  true,
		".yaml":  true,
		".yml":   true,
		".toml":  true,
		".ini":   true,
		".cfg":   true,
		".conf":  true,
		".sql":   true,
		".lua":   true,
		".vim":   true,
		".md":    true,
		".txt":   true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	return codeExtensions[ext]
}

// countLines đếm số dòng trong file
func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil
}

// scanDirectory quét thư mục và trả về danh sách file với số dòng
func scanDirectory(dirPath string, recursive bool) ([]FileInfo, error) {
	var fileInfos []FileInfo

	if recursive {
		// Quét đệ quy tất cả thư mục con
		err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("Lỗi khi truy cập %s: %v\n", path, err)
				return nil // Tiếp tục với các file khác
			}

			// Bỏ qua thư mục ẩn
			if d.IsDir() && strings.HasPrefix(d.Name(), ".") && path != dirPath {
				return filepath.SkipDir
			}

			// Chỉ xử lý file (không phải thư mục)
			if !d.IsDir() && isCodeFile(d.Name()) {
				lines, err := countLines(path)
				if err != nil {
					fmt.Printf("Lỗi khi đọc file %s: %v\n", path, err)
					return nil
				}

				fileInfos = append(fileInfos, FileInfo{
					Path:  path,
					Lines: lines,
				})
			}

			return nil
		})

		return fileInfos, err
	} else {
		// Chỉ quét thư mục hiện tại, không đệ quy
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() && isCodeFile(entry.Name()) {
				filePath := filepath.Join(dirPath, entry.Name())
				lines, err := countLines(filePath)
				if err != nil {
					fmt.Printf("Lỗi khi đọc file %s: %v\n", filePath, err)
					continue
				}

				fileInfos = append(fileInfos, FileInfo{
					Path:  filePath,
					Lines: lines,
				})
			}
		}
	}

	return fileInfos, nil
}

// printResults in kết quả ra màn hình
func printResults(fileInfos []FileInfo) {
	if len(fileInfos) == 0 {
		fmt.Println("Không tìm thấy file code nào trong thư mục.")
		return
	}

	// Sắp xếp theo số dòng tăng dần
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Lines < fileInfos[j].Lines
	})
	totalLines := 0
	totalCharacterLine := 80

	fmt.Printf("%-75s %10s\n", "ĐƯỜNG DẪN FILE", "SỐ DÒNG")
	fmt.Println(strings.Repeat("-", totalCharacterLine))

	for _, info := range fileInfos {
		length := totalCharacterLine - len(info.Path) - len(strconv.Itoa(info.Lines))

		fmt.Printf("%s %s %d\n", info.Path, strings.Repeat("-", length), info.Lines)
		totalLines += info.Lines
	}

	fmt.Println("")
	fmt.Printf("%s %d\n", "TỔNG CỘNG:", totalLines)
	fmt.Printf("Tổng số file: %d\n", len(fileInfos))
}

func main() {
	// Định nghĩa các flag
	var dirPath string
	var noRecursive bool
	var showHelp bool

	flag.StringVar(&dirPath, "dir", ".", "Đường dẫn thư mục cần quét (mặc định: thư mục hiện tại)")
	flag.StringVar(&dirPath, "d", ".", "Đường dẫn thư mục cần quét (viết tắt)")
	flag.BoolVar(&noRecursive, "no-recursive", false, "Chỉ quét thư mục hiện tại (không quét thư mục con)")
	flag.BoolVar(&noRecursive, "n", false, "Chỉ quét thư mục hiện tại (viết tắt)")
	flag.BoolVar(&showHelp, "help", false, "Hiển thị hướng dẫn sử dụng")
	flag.BoolVar(&showHelp, "h", false, "Hiển thị hướng dẫn sử dụng (viết tắt)")

	flag.Parse()

	if showHelp {
		fmt.Println("LINE COUNTER - Công cụ đếm dòng code")
		fmt.Println("\nCách sử dụng:")
		fmt.Println("  line-counter [options]")
		fmt.Println("\nTùy chọn:")
		fmt.Println("  -d, --dir          Đường dẫn thư mục cần quét (mặc định: thư mục hiện tại)")
		fmt.Println("  -n, --no-recursive Chỉ quét thư mục hiện tại (không quét thư mục con)")
		fmt.Println("  -h, --help         Hiển thị hướng dẫn sử dụng")
		fmt.Println("\nVí dụ:")
		fmt.Println("  line-counter")
		fmt.Println("  line-counter -d /path/to/project")
		fmt.Println("  line-counter -d /path/to/project -n")
		fmt.Println("\nCác loại file được hỗ trợ:")
		fmt.Println("  Go, JavaScript, TypeScript, Python, Java, C/C++, C#, PHP, Ruby,")
		fmt.Println("  Rust, Swift, Kotlin, Scala, R, Objective-C, Perl, Shell scripts,")
		fmt.Println("  HTML, CSS, Vue, React, XML, JSON, YAML, SQL, Lua, Markdown, v.v.")
		return
	}

	// Kiểm tra xem thư mục có tồn tại không
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("Lỗi: Thư mục '%s' không tồn tại.\n", dirPath)
		os.Exit(1)
	}

	fmt.Printf("Đang quét thư mục: %s\n", dirPath)
	if noRecursive {
		fmt.Println("Chế độ: Chỉ quét thư mục hiện tại")
	} else {
		fmt.Println("Chế độ: Quét đệ quy tất cả thư mục con")
	}
	fmt.Println()

	// Quét thư mục và đếm dòng
	recursive := !noRecursive
	fileInfos, err := scanDirectory(dirPath, recursive)
	if err != nil {
		fmt.Printf("Lỗi khi quét thư mục: %v\n", err)
		os.Exit(1)
	}

	// In kết quả
	printResults(fileInfos)
}

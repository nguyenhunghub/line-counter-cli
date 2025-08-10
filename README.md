# Line Counter (lco)

A fast and efficient command-line tool to count lines of code in your projects. Supports multiple programming languages and file types.

## Features

- Count lines in single files or entire directories
- Support for 30+ programming languages and file types
- Recursive directory scanning
- Sort results by file size (lines)
- Simple and fast execution

## Installation

### Prerequisites

- Go 1.16 or higher (for building from source)
- Git (optional, for cloning the repository)

### Using Make (Recommended)

```bash
# Clone the repository (optional if you already have the source)
git clone https://github.com/nguyenhunghub/line-counter.git
cd line-counter

# Build the binary
make b # or run: make go build -o lco.exe main.go

# Install to your PATH (optional)
sudo cp ./lco /usr/local/bin/
```


## Usage

### Basic Usage

Count lines in the current directory:

```bash
./lco
```

### Count Lines in a Specific Directory

```bash
./lco /path/to/directory
```

### Recursively Count Lines in Subdirectories

```bash
./lco -r /path/to/directory
```

### Show Help

```bash
./lco -h
```

### Command Line Options

```
Usage of ./lco:
  -r    Recursively scan subdirectories
  -h    Show help message
```

## Supported File Types

The tool supports counting lines in various file types including:

- Source code: .go, .js, .ts, .py, .java, .c, .cpp, .h, .hpp, .cs, .php, .rb, .rs, .swift, .kt, .scala, .r, .m, .pl
- Scripts: .sh, .bash, .zsh, .fish, .ps1
- Web: .html, .css, .scss, .sass, .less, .vue, .jsx, .tsx
- Data: .json, .yaml, .yml, .toml, .ini, .cfg, .conf, .sql
- Documentation: .md, .txt
- Other: .lua, .vim

## Examples

### Count lines in a Go project

```bash
# Count lines in the current directory (non-recursive)
./lco

# Count lines recursively in all subdirectories
./lco -r
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

<div align="center">

<picture>
  <img width="563" height="128" alt="tview LOGO" src="https://github.com/user-attachments/assets/1eee2c5c-a0c1-4c3b-a337-54b5f74dc708" />
</picture>

![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go)
![License MIT](https://img.shields.io/badge/License-MIT-green)

</div>

A fast, simple, and elegant terminal tool to visualize your folder structure as a beautiful tree with file icons and colors.

## ✨ Features

- 🎨 **Colorful output** with syntax highlighting for different file types
- 🏷️ **File icons** for better visual recognition
- 📊 **File size display** for quick size assessment
- 🎯 **Customizable depth** for focused directory exploration
- 🚫 **Ignore patterns** to exclude unwanted directories
- 💻 **Cross-platform** support (Windows, macOS, Linux)
- ⚡ **Fast and lightweight** written in Go

## 🖼️ Preview

![tview output example](https://github.com/user-attachments/assets/52b08699-b6cf-46f1-a0ff-6501dbb1a0b3)
*Example output of tview showing directory structure with icons and colors*

## 🚀 Installation

### Prerequisites
- Go 1.19 or higher

### Build from source

```bash
git clone https://github.com/sameer240704/tview
cd tview
go build -o tview main.go
```

## 📖 Usage

### Basic Usage
```bash
# Show current directory tree
tview

# Show specific directory tree
tview /path/to/directory

# Show with file sizes
tview --size

# Hide icons for minimal output
tview --icons=false
```

### Advanced Options
```bash
# Limit directory traversal depth
tview --depth=2

# Unlimited depth traversal
tview --depth=-1

# Ignore specific directories
tview --ignore="node_modules,.git,dist"

# Disable colored output
tview --color=false

# Show file sizes
tview --size

# Show version information
tview --version
```

## Command Line Options

The following table lists all available command line options for the tool:

| Option        | Short | Description                                      | Default  |
|---------------|-------|--------------------------------------------------|----------|
| `--depth`     |       | Maximum directory depth (`-1` for unlimited)    | `0`      |
| `--ignore`    |       | Comma-separated list of directories to ignore   | `""`     |
| `--color`     |       | Enable colored output                            | `true`   |
| `--icons`     |       | Hide file and folder icons                       | `false`  |
| `--size`      |       | Display file sizes                               | `false`  |
| `--version`   | `-v`  | Show version information                         | `false`  |
| `--help`      | `-?`  | Show help message                                | `false`  |

## 📝 LICENSE
This project is licensed under the MIT License - see the [LICENSE](https://github.com/sameer240704/tview?tab=MIT-1-ov-file) file for details.


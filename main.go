package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
	"tview/icons"
)

var (
    depth int
    ignore string
    useColor bool
	hideIcons bool
	hideFileSize bool
)

func init() {
    flag.IntVar(&depth, "depth", 0, "Max depth to traverse (-1 for infinite)")
    flag.StringVar(&ignore, "ignore", "", "Comma-separated list of ignored directories")
    flag.BoolVar(&useColor, "color", true, "Enable colored output")
	flag.BoolVar(&hideIcons, "icons", false, "Hide folder and file icons")
	flag.BoolVar(&hideFileSize, "size", false, "Display size of the corressponding file")
}

func main() {
    flag.Parse()
    root := "."
    if flag.NArg() > 0 {
        root = flag.Arg(0)
    }

	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error resolving path: ", err)
		return
	}

	rootName := filepath.Base(absRoot)

	rootIcon := " "
	if !hideIcons {
		rootIcon = ""
	}
    ignoreList := strings.Split(ignore, ",")

	var rootDisplay string
	if !hideIcons {
		rootDisplay = fmt.Sprintf("%s", rootName)
	} else {
		rootDisplay = fmt.Sprintf("%s %s", rootIcon, rootName)
	}

	fmt.Printf("%s\n", colorize(rootDisplay, "cyan", true))
    printTree(root, "", 0, ignoreList)
}

func shouldIgnore(path string, ignoreList []string) bool {
    for _, i := range ignoreList {
        if strings.TrimSpace(i) == filepath.Base(path) {
            return true
        }
    }
    return false
}

func colorize(text, color string, bold bool) string {
    if !useColor {
        return text
    }
    
    colors := map[string]string{
        "reset":  "\033[0m",
        "bold":   "\033[1m",
        "cyan":   "\033[36m",
        "blue":   "\033[34m",
        "green":  "\033[32m",
        "yellow": "\033[33m",
        "gray":   "\033[90m",
    }
    
    colorCode := colors[color]
    boldCode := ""
    if bold {
        boldCode = colors["bold"]
    }
    
    return boldCode + colorCode + text + colors["reset"]
}

func printTree(path, prefix string, level int, ignoreList []string) {
    if depth != -1 && level > depth {
        return
    }
    
    entries, err := os.ReadDir(path)
    if err != nil {
        fmt.Printf("%sError reading %s: %v\n", prefix, path, err)
        return
    }

    // Filter out ignored entries
    filteredEntries := make([]os.DirEntry, 0)
    for _, entry := range entries {
        if !shouldIgnore(entry.Name(), ignoreList) {
            filteredEntries = append(filteredEntries, entry)
        }
    }

    for i, entry := range filteredEntries {
        isLast := i == len(filteredEntries)-1
        
        // Curved connectors
        var connector, newPrefix string
        if isLast {
            connector = "╰── "
            newPrefix = prefix + "    "
        } else {
            connector = "├── "
            newPrefix = prefix + "│   "
        }
        
        // Color and styling
        displayName := entry.Name()

		var icon string
		if hideIcons {
			icon = icons.GetIcon(entry)
		} else {
			icon = ""
		}
        
        if entry.IsDir() {
            displayName = colorize(displayName, "cyan", true)
        } else {
            displayName = colorize(displayName, "blue", false)
        }

        if hideFileSize {
        	if !entry.IsDir() {
            	info, err := entry.Info()
	            if err == nil {
    	            size := formatSize(info.Size())
        	        sizeStr := colorize(fmt.Sprintf(" (%s)", size), "gray", false)
	                displayName += sizeStr
    	        }
	        }
		}
        
		if icon != "" {
            fmt.Printf("%s%s %s %s\n", 
                colorize(prefix, "gray", false),
                colorize(connector, "gray", false),
                icon,
                displayName)
        } else {
            fmt.Printf("%s%s %s\n", 
                colorize(prefix, "gray", false),
                colorize(connector, "gray", false),
                displayName)
        }
                
        if entry.IsDir() {
            printTree(filepath.Join(path, entry.Name()), newPrefix, level+1, ignoreList)
        }
    }
}

func formatSize(size int64) string {
    if size < 1024 {
        return fmt.Sprintf("%dB", size)
    } else if size < 1024*1024 {
        return fmt.Sprintf("%.1fKB", float64(size)/1024)
    } else if size < 1024*1024*1024 {
        return fmt.Sprintf("%.1fMB", float64(size)/(1024*1024))
    } else {
        return fmt.Sprintf("%.1fGB", float64(size)/(1024*1024*1024))
    }
}

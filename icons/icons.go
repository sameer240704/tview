package icons

import (
	"os"
	"path/filepath"
	"strings"
)

func GetIcon(entry os.DirEntry) string {
    if entry.IsDir() {
        return " "
    }
    
    ext := strings.ToLower(filepath.Ext(entry.Name()))
    switch ext {
	case ".go":
		return "󰟓 "
	case ".py", ".pyc":
		return " "
	case ".js", ".mjs":
		return " "
	case ".ts":
		return " "
	case ".tsx", ".jsx":
		return " "
	case ".java", ".class", ".jar":
		return " "
	case ".cpp":
		return " "
	case ".c":
		return " "
    case ".rs":
        return " "
	case ".css", ".scss":
		return " "
	case ".html":
		return " "
	case ".php":
		return " "
	case ".cs":
		return " "
	case ".swift":
		return " "
	case ".kt", ".kts":
		return " "
	case ".rb", ".rhtml":
		return " "
	case ".sql", ".mdf", ".ldf":
		return " "
	case ".csv":
		return " "
	case ".xlsx", ".xls":
		return "󱎏 "
	case ".pptx", ".ppt":
		return "󱎐 "
	case ".xml":
		return "󰗀 "
	case ".dockerignore":
		return " "
	case ".json":
		return " "
	case ".git", ".gitignore":
		return " "
	case ".config":
		return " "
	case ".md":
		return " "
	case ".txt":
		return "󰊄 "
	case ".yaml", ".yml":
		return " "
	case ".env", ".env.local":
		return " "
    case ".doc", ".docx":
        return "󰈙 "
    case ".jpg", ".jpeg", ".png", ".gif", ".svg", ".ico":
        return " "
    case ".mp4", ".avi", ".mov", ".mkv":
        return " "
    case ".mp3", ".wav", ".flac":
        return " "
	case ".zip":
		return " "
	case ".tar", ".gz", ".7z":
        return " "
    case ".pdf":
        return " "
    default:
        return " "
    }
}

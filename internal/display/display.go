package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"os"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

func DisplayFiles(files []fileinfo.FileInfo, options flags.Options) {
	if options.Long {
		displayLong(files)
	} else {
		displayShort(files)
	}
}

func displayShort(files []fileinfo.FileInfo) {
	for _, file := range files {
		fmt.Println(getColoredName(file.Name, file.Mode))
	}
}

func displayLong(files []fileinfo.FileInfo) {
	for _, file := range files {
		mode := file.Mode.String()
		links := 1
		size := file.Size
		timeStr := time.Unix(file.ModTime, 0).Format("Jan _2 15:04")
		name := getColoredName(file.Name, file.Mode)

		fmt.Printf("%s %2d %s %s %6d %s %s\n", mode, links, file.User, file.Group, size, timeStr, name)
	}
}

func getColoredName(name string, mode os.FileMode) string {
	switch {
	case mode.IsDir():
		return colorBlue + name + colorReset
	case mode&os.ModeSymlink != 0:
		return colorCyan + name + colorReset
	case mode&os.ModeNamedPipe != 0:
		return colorYellow + name + colorReset
	case mode&os.ModeSocket != 0:
		return colorGreen + name + colorReset
	case mode&os.ModeDevice != 0:
		return colorYellow + name + colorReset
	case mode&0111 != 0:
		return colorRed + name + colorReset
	default:
		return name
	}
}

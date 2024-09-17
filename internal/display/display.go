package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"os"
	"time"
)

const (
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
	colorGreen = "\033[32m"
	colorCyan  = "\033[36m"
)

func DisplayFiles(files []fileinfo.FileInfo, totalBlocks int64, options flags.Options, isDir bool) {
	if options.Long {
		displayLong(files, totalBlocks, isDir)
	} else {
		displayShort(files)
	}
}

func displayShort(files []fileinfo.FileInfo) {
	for _, file := range files {
		fmt.Println(getColoredName(file))
	}
}

func displayLong(files []fileinfo.FileInfo, totalBlocks int64, isDir bool) {
	if isDir {
		fmt.Printf("total %d\n", totalBlocks/2) 
	}

	for _, file := range files {
		mode := file.Mode.String()
		links := file.NumLinks
		size := file.Size
		timeStr := time.Unix(file.ModTime, 0).Format("Jan _2 15:04")
		name := getColoredName(file)

		fmt.Printf("%s %2d %s %s %6d %s %s\n", mode, links, file.User, file.Group, size, timeStr, name)
	}
}

func getColoredName(file fileinfo.FileInfo) string {
	switch {
	case file.IsDir:
		return colorBlue + file.Name + colorReset
	case file.Mode&0111 != 0:
		return colorGreen + file.Name + colorReset
	case file.Mode&os.ModeSymlink != 0:
		return colorCyan + file.Name + colorReset
	default:
		return file.Name
	}
}

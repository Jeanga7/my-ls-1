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

func DisplayFiles(files []fileinfo.FileInfo, options flags.Options) {
	if options.Long {
		displayLong(files)
	} else {
		displayShort(files)
	}
}

func displayShort(files []fileinfo.FileInfo) {
	for _, file := range files {
		if file.Name != "total" {
			fmt.Println(getColoredName(file))
		}
	}
}

func displayLong(files []fileinfo.FileInfo) {
	var totalBlocks int64
	var regularFiles []fileinfo.FileInfo

	for _, file := range files {
		if file.Name == "total" {
			totalBlocks = file.BlockSize
		} else {
			regularFiles = append(regularFiles, file)
		}
	}

	fmt.Printf("total %d\n", totalBlocks)

	for _, file := range regularFiles {
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

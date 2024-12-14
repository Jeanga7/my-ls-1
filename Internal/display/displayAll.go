package display

import (
	"fmt"
	"my_ls/Internal/fileinfo"
)

func DisplayAll(files []fileinfo.FileInfo) {
	for i, file := range files {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(getColoredName(file))
	}
	fmt.Println()
}

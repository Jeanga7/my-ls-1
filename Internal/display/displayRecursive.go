package display

import (
	"fmt"
	"my_ls/Internal/fileinfo"
)

func DisplayRecursive(files []fileinfo.FileInfo) {
	for _, file := range files {
		fmt.Println(getColoredName(file))
	}
}


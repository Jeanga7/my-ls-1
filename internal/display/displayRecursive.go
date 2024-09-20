package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
)

func DisplayRecursive(files []fileinfo.FileInfo) {
	for _, file := range files {
		fmt.Println(getColoredName(file))
	}
}


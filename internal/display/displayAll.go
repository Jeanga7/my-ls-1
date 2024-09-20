package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
)

func DisplayAll(files []fileinfo.FileInfo) {
	for _, file := range files {
		fmt.Println(getColoredName(file))
	}
}
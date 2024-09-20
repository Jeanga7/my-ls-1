package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
)

func DisplayDefault(files []fileinfo.FileInfo) {
	for _, file := range files {
		if file.Name[0] != '.' {
			fmt.Println(getColoredName(file))
		}
	}
}
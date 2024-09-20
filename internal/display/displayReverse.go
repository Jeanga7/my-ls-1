package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
)

func DisplayReverse(files []fileinfo.FileInfo) {
	for i := len(files) - 1; i >= 0; i-- {
		fmt.Println(getColoredName(files[i]))
	}
}
package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
)

func DisplayReverse(files []fileinfo.FileInfo) {
	for i := len(files) - 1; i >= 0; i-- {
		if i < len(files) - 1 {
			fmt.Print(" ")
		}
		fmt.Print(getColoredName(files[i]))
	}
	fmt.Println() 
}
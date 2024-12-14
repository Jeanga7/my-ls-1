package display

import (
	"fmt"
	"my_ls/Internal/fileinfo"
	"sort"
)

func DisplayTime(files []fileinfo.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime > files[j].ModTime
	})

	for i, file := range files {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(getColoredName(file))
	}
	fmt.Println() 
}
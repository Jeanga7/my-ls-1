package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
	"sort"
)

func DisplayTime(files []fileinfo.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime > files[j].ModTime
	})

	for _, file := range files {
		fmt.Println(getColoredName(file))
	}
}
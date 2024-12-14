package sort

import (
	"my_ls/Internal/fileinfo"
	"my_ls/Internal/flags"
	"strings"
)

func SortFiles(files []fileinfo.FileInfo, options flags.Options) {
	bubbleSort(files, func(i, j int) bool {
		if options.Time {
			if files[i].ModTime != files[j].ModTime {
				if options.Reverse {
					return files[i].ModTime < files[j].ModTime
				}
				return files[i].ModTime > files[j].ModTime
			}
		}

		iLower := strings.ToLower(files[i].Name)
		jLower := strings.ToLower(files[j].Name)

		iLower = strings.TrimPrefix(iLower, ".")
        jLower = strings.TrimPrefix(jLower, ".")

		if options.Reverse {
			return iLower < jLower
		}
		return iLower < jLower
	})
}

func bubbleSort(files []fileinfo.FileInfo, less func(i, j int) bool) {
	n := len(files)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if less(j+1, j) {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}


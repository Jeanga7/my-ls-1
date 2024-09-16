package sort

import (
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"sort"
)

func SortFiles(files []fileinfo.FileInfo, options flags.Options) {
	sort.Slice(files, func(i, j int) bool {
		if options.Time {
			if files[i].ModTime != files[j].ModTime {
				if options.Reverse {
					return files[i].ModTime < files[j].ModTime
				}
				return files[i].ModTime > files[j].ModTime
			}
		}

		if options.Reverse {
			return files[i].Name > files[j].Name
		}
		return files[i].Name < files[j].Name
	})
}

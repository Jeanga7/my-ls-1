package display

import (
	"fmt"
	"my_ls/internal/fileinfo"
	"os"
	"time"
)

func DisplayLong(files []fileinfo.FileInfo, totalBlocks int64, isDir bool) {
	if isDir {
		fmt.Printf("total %d\n", totalBlocks/2)
	}

	for _, file := range files {
		mode := file.Mode.String()
		if file.Mode&os.ModeSymlink != 0 {
			mode = "l" + mode[1:]
		}
		links := file.NumLinks
		size := file.Size
		timeStr := formatTime(file.ModTime)
		var name string
		if file.Mode&os.ModeSymlink != 0{
			name = ColorBold + ColorCyan + file.Name + ColorReset
		}else{
			name = getColoredName(file)
		}

		fmt.Printf("%s %2d %s %s %6d %s %s", mode, links, file.User, file.Group, size, timeStr, name)

		if file.Mode&os.ModeSymlink != 0 {
			target, err := os.Readlink(file.Path)
			if err == nil {
				fmt.Printf(" -> %s", ColorBold + ColorGreen + target + ColorReset)
			}
		}
		fmt.Println()
	}
}

func formatTime(modTime int64) string {
	t := time.Unix(modTime, 0)
	now := time.Now()
	sixMonthsAgo := now.AddDate(0, -6, 0)

	if t.After(sixMonthsAgo) {
		return t.Format("Jan _2 15:04")
	}
	return t.Format("Jan _2  2006")
}
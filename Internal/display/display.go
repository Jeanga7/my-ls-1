package display

import (
	"my_ls/Internal/fileinfo"
	"my_ls/Internal/flags"
	"os"
)

const (
	ColorReset = "\033[0m"
	ColorBlue  = "\033[34m"
	ColorGreen = "\033[32m"
	ColorCyan  = "\033[36m"
	ColorBold  = "\033[1m"
)

func DisplayFiles(files []fileinfo.FileInfo, totalBlocks int64, options flags.Options, isDir bool) {
	
	if options.Long {
		DisplayLong(files, totalBlocks, isDir)
	} else if options.Recursive {
		DisplayRecursive(files)
	} else if options.Reverse {
		DisplayReverse(files)
	} else if options.Time {
		DisplayTime(files)
	} else if options.All {
		DisplayAll(files)
	} else {
		DisplayDefault(files)
	}
}

func getColoredName(file fileinfo.FileInfo) string {
	switch {
	case file.IsDir:
		return ColorBold + ColorBlue + file.Name + ColorReset
	case file.Mode&0111 != 0:
		return ColorBold + ColorGreen + file.Name + ColorReset
	case file.Mode&os.ModeSymlink != 0:
		return ColorBold + ColorCyan + file.Name + ColorReset
	default:
		return file.Name
	}
}

package main

import (
	"fmt"
	"my_ls/internal/display"
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"my_ls/internal/sort"
	"os"
	"strings"
)

func parseFlags(args []string) ([]string, flags.Options) {
	options := flags.Options{}
	var paths []string

	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-") {
			for _, flag := range arg[1:] {
				switch flag {
				case 'l':
					options.Long = true
				case 'R':
					options.Recursive = true
				case 'a':
					options.All = true
				case 'r':
					options.Reverse = true
				case 't':
					options.Time = true
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	return paths, options
}

func main() {
	paths, options := parseFlags(os.Args)

	for _, path := range paths {
		err := processPath(path, options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}

func processPath(path string, options flags.Options) error {
	files, err := fileinfo.GetFileInfo(path, options)
	if err != nil {
		return err
	}

	sort.SortFiles(files, options)

	display.DisplayFiles(files, options)

	return nil
}

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

	if len(paths) > 1 {
		for i, path := range paths {
			if i > 0 {
				fmt.Println()
			}
			// fmt.Printf("%s:\n", path)
			err := processPath(path, options, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
		}
	} else {
		err := processPath(paths[0], options, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}

func processPath(path string, options flags.Options, forceDir bool) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		files, totalBlocks, err := fileinfo.GetFileInfo(path, options)
		if err != nil {
			return err
		}

		sort.SortFiles(files, options)
		display.DisplayFiles(files, totalBlocks, options, true)
	} else {
		file, err := fileinfo.GetSingleFileInfo(path)
		if err != nil {
			return err
		}
		display.DisplayFiles([]fileinfo.FileInfo{file}, file.BlockSize, options, forceDir)
	}

	return nil
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

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

	for i, path := range paths {
		if i > 0 {
			fmt.Println()
		}
		err := processPath(path, options, len(paths) > 1 || options.Recursive)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}

func processPath(path string, options flags.Options, printHeader bool) error {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		if printHeader {
			fmt.Printf("%s:\n", path)
		}
		err = processDirectory(path, options)
		if err != nil {
			return err
		}
		if options.Recursive {
			err = processRecursive(path, options)
		}
	} else {
		file, err := fileinfo.GetSingleFileInfo(path)
		if err != nil {
			return err
		}
		display.DisplayFiles([]fileinfo.FileInfo{file}, file.BlockSize, options, false)
	}

	return nil
}

func processDirectory(path string, options flags.Options) error {
	files, totalBlocks, err := fileinfo.GetFileInfo(path, options)
	if err != nil {
		return err
	}

	sort.SortFiles(files, options)
	display.DisplayFiles(files, totalBlocks, options, true)
	return nil
}

func processRecursive(path string, options flags.Options) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subpath := path + "/" + entry.Name()
			if !options.All && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			fmt.Printf("\n./%s:\n", subpath)
			err = processDirectory(subpath, options)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error processing %s: %v\n", subpath, err)
			}
			err = processRecursive(subpath, options)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error processing %s: %v\n", subpath, err)
			}
		}
	}
	return nil
}
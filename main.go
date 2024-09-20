package main

import (
	"fmt"
	"my_ls/internal/display"
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"my_ls/internal/sort"
	"my_ls/internal/utils"
	"os"
	"strings"
)

func main() {
	paths, options := utils.ParseFlags(os.Args)

	if len(paths) == 0 {
		paths = []string{"."}
	}

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
	} else {
		file, err := fileinfo.GetSingleFileInfo(path)
		if err != nil {
			return err
		}
		display.DisplayFiles([]fileinfo.FileInfo{file}, file.BlockSize, options, false)
	}
	return err
}

func processDirectory(path string, options flags.Options) error {
	files, _, err := fileinfo.GetFileInfo(path, options)
	if err != nil {
		return err
	}

	if options.All {
		dotFile, _ := fileinfo.GetSingleFileInfo(path + "/.")
		dotDotFile, _ := fileinfo.GetSingleFileInfo(path + "/..")
		files = append([]fileinfo.FileInfo{dotFile, dotDotFile}, files...)
	} else {
		visibleFiles := []fileinfo.FileInfo{}
		for _, file := range files {
			if file.Name[0] != '.' {
				visibleFiles = append(visibleFiles, file)
			}
		}
		files = visibleFiles
	}

	sort.SortFiles(files, options)
	display.DisplayFiles(files, 0, options, true)

	return nil
}

func ProcessRecursive(path string, options flags.Options) error {
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
			err = ProcessRecursive(subpath, options)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error processing %s: %v\n", subpath, err)
			}
		}
	}
	return nil
}

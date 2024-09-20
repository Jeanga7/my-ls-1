package utils

import (
	"my_ls/internal/flags"
	"os"
	"strings"
	"syscall"
)

func ParseFlags(args []string) ([]string, flags.Options) {
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

func GetLinkTarget(path string) (string, error) {
	target, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	if !isAbs(target) {
		target = join(dir(path), target)
	}
	return target, nil
}

func CountHardLinks(info os.FileInfo) int {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return int(stat.Nlink)
	}
	return 1
}

func isAbs(path string) bool {
	return strings.HasPrefix(path, "/")
}

func join(elem ...string) string {
	return strings.Join(elem, "/")
}

func dir(path string) string {
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return "."
	}
	if i == 0 {
		return "/"
	}
	return path[:i]
}

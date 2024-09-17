package utils

import (
	"os"
	"strings"
	"syscall"
)

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
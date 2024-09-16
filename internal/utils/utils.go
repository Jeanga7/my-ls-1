package utils

import (
	"os"
	"path/filepath"
	"runtime"
	//"syscall"
)

// GetLinkTarget retourne la cible d'un lien symbolique
func GetLinkTarget(path string) (string, error) {
	target, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	if !filepath.IsAbs(target) {
		target = filepath.Join(filepath.Dir(path), target)
	}
	return target, nil
}

func CountHardLinks(info os.FileInfo) int {
	// Pour Windows, retourne 1 par défaut
	if runtime.GOOS == "windows" {
		return 1
	}

	// Pour Unix/Linux, utilise syscall.Stat_t
	//if stat, ok := info.Sys().(*syscall.Stat_t); ok {
	//	return int(stat.Nlink)
	//}
	return 1
}

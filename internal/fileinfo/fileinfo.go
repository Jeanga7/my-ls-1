package fileinfo

import (
	"my_ls/internal/flags"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	//"syscall"
)

type FileInfo struct {
	Name    string
	Path    string
	IsDir   bool
	Size    int64
	Mode    os.FileMode
	ModTime int64
	User    string
	Group   string
}

func GetFileInfo(path string, options flags.Options) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignorer les fichiers cachés si l'option -a n'est pas activée
		if !options.All && info.Name()[0] == '.' && info.Name() != "." {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		file := FileInfo{
			Name:    info.Name(),
			Path:    path,
			IsDir:   info.IsDir(),
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime().Unix(),
		}

		// Ajouter les informations sur l'utilisateur et le groupe si l'option -l est activée
		if options.Long {
			uid, gid := getFileOwner(info)
			file.User = getUserName(uid)
			file.Group = getGroupName(gid)
		}

		files = append(files, file)

		if !options.Recursive && info.IsDir() && path != "." {
			return filepath.SkipDir
		}

		return nil
	})

	return files, err
}

// getFileOwner retourne l'UID et le GID sous Unix/Linux ou des valeurs par défaut sous Windows
func getFileOwner(info os.FileInfo) (uint32, uint32) {
	if runtime.GOOS == "windows" {
		return 0, 0
	}

	// Pour Unix/Linux, utilise syscall.Stat_t
	//if stat, ok := info.Sys().(*syscall.Stat_t); ok {
	//	return stat.Uid, stat.Gid
	//}

	// Retourne 0, 0 si aucune information n'est disponible
	return 0, 0
}

func getUserName(uid uint32) string {
	u, err := user.LookupId(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		return strconv.FormatUint(uint64(uid), 10)
	}
	return u.Username
}

func getGroupName(gid uint32) string {
	g, err := user.LookupGroupId(strconv.FormatUint(uint64(gid), 10))
	if err != nil {
		return strconv.FormatUint(uint64(gid), 10)
	}
	return g.Name
}

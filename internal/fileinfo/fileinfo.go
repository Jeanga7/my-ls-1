package fileinfo

import (
	"my_ls/internal/flags"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
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

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !options.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := path + "/" + entry.Name()
		file := FileInfo{
			Name:    entry.Name(),
			Path:    fullPath,
			IsDir:   entry.IsDir(),
			Size:    entry.Size(),
			Mode:    entry.Mode(),
			ModTime: entry.ModTime().Unix(),
		}

		if options.Long {
			uid, gid := getFileOwner(entry)
			file.User = getUserName(uid)
			file.Group = getGroupName(gid)
		}

		files = append(files, file)
	}

	return files, nil
}

func GetSingleFileInfo(path string) (FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return FileInfo{}, err
	}

	file := FileInfo{
		Name:    info.Name(),
		Path:    path,
		IsDir:   info.IsDir(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime().Unix(),
	}

	uid, gid := getFileOwner(info)
	file.User = getUserName(uid)
	file.Group = getGroupName(gid)

	return file, nil
}

func walkDir(root string, fn func(string, os.FileInfo, error) error) error {
	info, err := os.Lstat(root)
	if err != nil {
		return fn(root, nil, err)
	}
	return walkDirRecur(root, info, fn)
}

func walkDirRecur(path string, info os.FileInfo, fn func(string, os.FileInfo, error) error) error {
	if !info.IsDir() {
		return fn(path, info, nil)
	}

	err := fn(path, info, nil)
	if err != nil {
		if err == os.ErrPermission {
			return nil
		}
		return err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return fn(path, info, err)
	}

	for _, entry := range entries {
		subInfo, err := entry.Info()
		if err != nil {
			if err := fn(strings.Join([]string{path, entry.Name()}, "/"), nil, err); err != nil && err != os.ErrPermission {
				return err
			}
		} else {
			err = walkDirRecur(strings.Join([]string{path, entry.Name()}, "/"), subInfo, fn)
			if err != nil {
				if err != os.ErrPermission {
					return err
				}
			}
		}
	}

	return nil
}

func getFileOwner(info os.FileInfo) (uint32, uint32) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return stat.Uid, stat.Gid
	}
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

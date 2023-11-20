package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(path.Dir(b))
}

func Path(p string) string {
	root := RootDir()
	return path.Join(root, p)
}

func CreateFileWithRootPath(path string, name string) (*os.File, error) {

	dir := Path(path)

	d, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
		d, err = os.Stat(dir)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	if d.IsDir() {
		file, err := os.Create(dir + "/" + name)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return nil, err
	}

}

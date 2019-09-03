package types

import (
	"errors"
	"mime"
	"net/http"
)

type MultiplexFileSystem struct {
	fsList []http.FileSystem
}

func init() {
	// https://github.com/labstack/echo/issues/1038#issuecomment-410294904
	mime.AddExtensionType(".js", "application/javascript")
}

func (ffs *MultiplexFileSystem) Add(fs http.FileSystem) {
	ffs.fsList = append(ffs.fsList, fs)
}

func (ffs MultiplexFileSystem) Open(name string) (http.File, error) {
	for _, item := range ffs.fsList {
		file, err := item.Open(name)
		if err != nil {
			continue
		}
		return file, nil
	}
	return nil, errors.New(name + " not found.")
}

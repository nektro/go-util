package types

import (
	"errors"
	"mime"
	"net/http"
)

type MultiplexFileSystem struct {
	FSList []http.FileSystem
}

func init() {
	// https://github.com/labstack/echo/issues/1038#issuecomment-410294904
	mime.AddExtensionType(".js", "application/javascript")
}

func (ffs MultiplexFileSystem) Open(name string) (http.File, error) {
	for _, item := range ffs.FSList {
		file, err := item.Open(name)
		if err != nil {
			continue
		}
		return file, nil
	}
	return nil, errors.New(name + " not found.")
}

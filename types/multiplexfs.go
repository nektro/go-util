package types

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/valyala/fasthttp"
)

type MultiplexFileSystem struct {
	FSList []http.FileSystem
}

func init() {
	// https://github.com/labstack/echo/issues/1038#issuecomment-410294904
	mime.AddExtensionType(".js", "application/javascript")
}

func (ffs MultiplexFileSystem) Open(name string) (http.File, error) {
	var errr error
	for _, item := range ffs.FSList {
		file, err := item.Open(name)
		if file != nil {
			return file, nil
		}
		errr = err
	}
	return nil, errr
}

func (ffs MultiplexFileSystem) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fpath := string(ctx.Request.RequestURI())
	if strings.HasSuffix(fpath, "/") {
		fpath += "index.html"
	}
	file, err := ffs.Open(fpath)
	if err == nil {
		bytes, _ := ioutil.ReadAll(file)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType(mime.TypeByExtension(path.Ext(fpath)))
		ctx.SetBody(bytes)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

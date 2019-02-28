package util

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	fasthttpHandlers = map[string]func(ctx *fasthttp.RequestCtx){}
)

func Log(message ...interface{}) {
	fmt.Print("[" + GetIsoDateTime() + "] ")
	fmt.Println(message...)
}

func GetIsoDateTime() string {
	vil := time.Now().UTC().String()
	return vil[0:19]
}

func FasthttpAddHandler(path string, handle func(ctx *fasthttp.RequestCtx)) {
	fasthttpHandlers[path] = handle
}

func FasthttpHandle(path string, ctx *fasthttp.RequestCtx) bool {
	for k, v := range fasthttpHandlers {
		if k == path {
			v(ctx)
			return true
		}
	}
	return false
}

package fasthttproute

import (
	"github.com/valyala/fasthttp"
)

var notFoundHandler fasthttp.RequestHandler

func DefaultHandler(ctx *fasthttp.RequestCtx) {
	sPath := Bytes2str(ctx.Path())

	if handler, ok := mapping[sPath]; ok {
		handler(ctx)
	} else {
		if notFoundHandler != nil {
			notFoundHandler(ctx)
		} else {
			defaultNotFoundHandler(ctx)
		}
	}
}

func defaultRootHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetBodyString("/")
}

func defaultFaviconHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "image/x-icon")
	ctx.SetStatusCode(204)
}

func defaultNotFoundHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(404)
	ctx.Response.SetBody([]byte("404, Not found"))
}

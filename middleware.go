package fasthttproute

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/valyala/fasthttp"
	"skipmeal/Helper"
)

type Middelware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Chain(handler fasthttp.RequestHandler, middelwares ...Middelware) fasthttp.RequestHandler {
	for _, m := range middelwares {
		handler = m(handler)
	}
	return handler
}

func Etag(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		body := ctx.Response.Body()
		bodySum := md5.Sum(body)
		var strEtag [md5.Size * 2]byte
		hex.Encode(strEtag[:], bodySum[:])
		if bytes.Equal(strEtag[:], ctx.Request.Header.Peek("If-None-Match")) {
			ctx.Response.Header.SetStatusCode(304)
			ctx.Response.SetBodyString("")
			return
		}
		ctx.Response.Header.Set("Etag", Helper.Bytes2str(strEtag[:]))
	}
}

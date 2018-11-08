package fasthttproute

import (
	"github.com/valyala/fasthttp"
)

type Middelware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Chain(handler fasthttp.RequestHandler, middelwares ...Middelware) fasthttp.RequestHandler {
	for _, m := range middelwares {
		handler = m(handler)
	}
	return  handler
}

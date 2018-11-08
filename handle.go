package fasthttproute

import (
	"github.com/valyala/fasthttp"
	"log"
)

var mapping = map[string]func(*fasthttp.RequestCtx){
	"/":            defaultRootHandler,
	"/favicon.ico": defaultFaviconHandler,
}

func Handle(path string, handler fasthttp.RequestHandler) {
	mapping[path] = handler
}

func HandleNotFound(handler fasthttp.RequestHandler) {
	notFoundHandler = handler
}

func ServeFasthttp(addr string, compress bool, handler fasthttp.RequestHandler) {
	if compress {
		handler = fasthttp.CompressHandler(handler)
	}
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

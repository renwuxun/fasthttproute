#### fasthttproute
应用于fasthttp的路由

#### 使用
>fasthttproute.Handle("/hello", func(ctx *fasthttp.RequestCtx) {
   ctx.SetBodyString("hello!")
})
compress := false
fasthttproute.ServeFasthttp(":80", compress, fasthttproute.DefaultHandler)

#### LICENSE
MIT License
#### fasthttproute
应用于fasthttp的路由

#### 使用
```go
fasthttproute.Handle("/hello", func(ctx *fasthttp.RequestCtx) {
   ctx.SetBodyString("hi,fasthttproute!")
})

fasthttproute.ServeFasthttp(":80", fasthttproute.DefaultHandler)
```
#### LICENSE
MIT License

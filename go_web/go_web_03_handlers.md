# Go语言五个内置的Handler

## NotFoundHandler 用来返回404 not found

+ func NotFoundHandler() Handler
+ 返回一个 handler，它给每个请求的响应都是“404 page not found

## RedirectHandler 重定向

+ func RedirectHandler(url string, code int) Handler
+ 返回一个handler， 它把每个请求给定的状态码跳转到指定的URL
  + url，要跳转到的 URL
  + code，跳转的状态码（3xx），常见的：StatusMovedPermanently、StatusFound 或 StatusSeeOther 等

## StripPrefix

+ func StripPrefix(prefix string, h handler)Handler
+ 返回一个handler，他从请丢URL中去掉指定的前缀，然后再调用另一个handler
  + 若请求的URL与提供的前缀不符，那么404
+ 略像中间件
  + prefix, URL将要被移除的字符串前缀
  + h 是一个handler，在移除之后这个handler会收到请求
+ 修饰了另一个handler

## TimeoutHandler

+ func TimeoutHandler(h Handler, dt time.Duration, msg string)Handler
+ 返回一个handler，它用来在指定时间内运行传入的h
+ 也相当与一个修饰器
  + h，将要被修饰的handler
  + dt，第一个handler允许的处理时间
  + msg, 如果超时，把msg返回给请求，表示响应时间过长

## FileServer

+ func FileServer(root FileSystem)Handler
+ 返回一个handler，使用基于root的文件系统来响应请求

```golang
type FileSystem interface {
    Open(name string) (File, error)
}
```

+ 使用时需要用到操作系统的文件系统，所以还需要委托给：
+ type Dir string
+ func (d Dir) Open(name string) (File, error)

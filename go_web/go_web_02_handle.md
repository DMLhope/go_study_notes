# class 02 handle请求

## 处理（handle）请求

http.Server is a struct

+ Addr 字段表示网络地址
+ Handler 字段 nill表示使用 DefaultServeMux（默认路由）
+ ListenAndServe 函数

Handler 是一个接口（interface）

Handler 定义了一个方法 ServeHTTP()包含

1. HTTPResponeWriter
2. 指向Request这个struct的指针

## DefaultServeMux

 它是一个Multiplexer（多路复用器） （具体就是用来做路由）
 它也是个Handler

## 多个Handler - http.Handler

不指定 Server struct 里面的 Handler 字段
可以使用 http.Handler 将某个Handler 附加（注册）到 DefaultServemux
  http 包有一个Handle函数
  ServerMux struct 也有一个Handler方法
如果掉用http.Handle 实际上调用的事DefaultServeMux上的Handle方法
DefaultServeMux 就是 ServeMux 的指针变量

## Handle 函数 - http.HandleFunc

Handle 函数 就是那些行为与handler相似的函数：
Handle 函数的签名与ServeHTTP方法的签名一样，接收：
  一个http.ResponseWriter
  一个指向http.Request的指针

## http.HandleFunc原理

Go 有一个函数类型： HandlerFunc。
可以将某个具有适当签名的函数f，适配成为一个Handler，而这个Handler具有方法f。

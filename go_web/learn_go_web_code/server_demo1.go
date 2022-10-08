package main

import "net/http"

// class 02 handle请求
//处理（handle）请求
//http.Server is a struct
// Addr 字段表示网络地址
// Handler 字段 nill表示使用 DefaultServeMux（默认路由）
// ListenAndServe 函数

// Handler 是一个接口（interface）
// Handler 定义了一个方法 ServeHTTP()包含
// 	1. HTTPResponeWriter
// 	2. 指向Request这个struct的指针

// DefaultServeMux
//  它是一个Multiplexer（多路复用器） （具体就是用来做路由）
//  它也是个Handler

// 多个Handler - http.Handler
// 不指定 Server struct 里面的 Handler 字段
// 可以使用 http.Handler 将某个Handler 附加（注册）到 DefaultServemux
//   http 包有一个Handle函数
//   ServerMux struct 也有一个Handler方法
// 如果掉用http.Handle 实际上调用的事DefaultServeMux上的Handle方法
//	 DefaultServeMux 就是 ServeMux 的指针变量

// Handle 函数 - http.HandleFunc
// Handle 函数 就是那些行为与handler相似的函数：
// Handle 函数的签名与ServeHTTP方法的签名一样，接收：
//   一个http.ResponseWriter
//   一个指向http.Request的指针

// http.HandleFunc原理
// Go 有一个函数类型： HandlerFunc。
// 可以将某个具有适当签名的函数f，适配成为一个Handler，而这个Handler具有方法f。

//==========================================================================================
// type myHandler struct{}

// func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("web test 2"))
// }

type testHandler struct{}

func (m *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("web test 2"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About !!!"))
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcome ~~~"))
}

func main2() {

	// 方法1
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: nil,
	// }
	// server.ListenAndServe()
	// 方法2
	//http.ListenAndServe("localhost:8080",nil)

	//方法1和方法2等价

	//上面两种方法创建的web服务器无法走https
	//想走得使用配套的
	//http.ListenServeTLS() 和 server.ListenServeTLS()

	// 这里就是使用了自己定义的Handler来响应web请求
	// 但是这样所有的请求就都会对应到这一个handler上
	// mh := myHandler{}
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: &mh,
	// }
	// server.ListenAndServe()

	test := testHandler{}
	about := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.Handle("/test", &test)
	http.Handle("/about", &about)

	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Home ~~~"))
	})
	http.HandleFunc("/welcome", welcome)

	//这里就是用http.HandlerFunc将welcome函数转换成一个handler
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()

}

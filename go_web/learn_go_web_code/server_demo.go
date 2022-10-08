package main

import "net/http"

func main1() {
	//调用handlefunc接收两个参数
	//1. 路由地址“/”表示根地址，即对所有的请求进行响应
	//2. 函数 接收两个参数
	// http.ResponseWriter 写响应
	// *http.Request指针表示的是传入请求的所有的信息
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test web"))
	})

	//设置web服务器
	http.ListenAndServe("localhost:8080", nil) //nill表示使用DefaultServeMux（路由）
}

//处理（handle）请求
//http.Server is a struct
// Addr 字段表示网络地址
// Handler 字段 nill表示使用DefaultServeMux（默认路由）
// ListenAndServe 函数

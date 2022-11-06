# ResponseWriter

## ResponseWriter 接口

+ 从服务器向客户端返回响应需要使用 ResponseWriter
+ ResponseWriter 是一个接口，handler 用它来返回响应
+ 真正支撑 ResponseWriter 的幕后 struct 是非导出的 http.response

## 问题

+ 为什么 Handler 的 ServeHTTP(w ResponseWriter, r *Request) ，只有一个是指针类型？而 w 是按值传递的吗？
+ 答：
  1. ResponseWriter接口代表了一个指针，是指向response这个struct的指针，所以ResponseWriter也是一个指针
  2. w也是按引用进行传递

## 写入到 ResponseWriter

+ Write 方法接收一个 byte 切片作为参数，然后把它写入到 HTTP 响应的 Body 里面。
+ 如果在 Write 方法被调用时，header 里面没有设定 content type，那么数据的前 512 字节就会被用来检测 content type

### 代码示例

```go
package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
 str := `
 <html>
 <head>
    <title>test</title>
 </head>
 <body>
  <h1>Go web</h1>
 </body>
 </html>`
 w.Write([]byte(str))
}

func main() {
 server := http.Server{
  Addr: "localhost:9000",
 }
 http.HandleFunc("/write", writeExample)
 server.ListenAndServe()
}
```

## WriteHeader 方法

+ WriteHeader 方法接收一个整数类型（HTTP 状态码）作为参数，并把它作为 HTTP 响应的状态码返回
+ 如果该方法没有显式调用，那么在第一次调用 Write 方法前，会隐式的调用 WriteHeader(http.StatusOK)
  + 所以 WriteHeader 主要用来发送错误类的 HTTP 状态码
+ 调用完 WriteHeader 方法之后，仍然可以写入到 ResponseWriter，但无法再修改 header 了

### 代码示例

```go
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service,try next door")
}

func main() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}
```

## Header 方法

+ Header 方法返回 headers 的 map，可以进行修改
+ 修改后的 headers 将会体现在返回给客户端的 HTTP 响应里
  + 设置header里的值一定要在调用WriteHeader之前

### 代码示例

```go
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "http://www.google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}
```

### 代码示例

```go
type Post struct {
	User   string
	Theads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:   "test",
		Theads: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
```

## 内置的 Response

+ NotFound 函数，包装一个 404 状态码和一个额外的信息
+ ServeFile 函数，从文件系统提供文件，返回给请求者
+ ServeContent 函数，它可以把实现了 io.ReadSeeker 接口的任何东西里面的内容返回给请求者
  + 还可以处理 Range 请求（范围请求），如果只请求了资源的一部分内容，那么 ServeContent 就可以如此响应。而 ServeFile 或 io.Copy 则不行。
+ Redirect 函数，告诉客户端重定向到另一个 URL

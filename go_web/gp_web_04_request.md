# 请求 Request

## HTTP请求

+ Request
+ URL
+ Header
+ Body

## HTTP 消息

HTTP Request 和 HTTP Response （请求和响应）
他们具有相同的结构：

+ 请求（响应）行
+ 0个或多个Header
+ 空行
+ 可选的消息体（Body）
例子：

```html
GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(空行)
```

net/http包提供了用于表示HTTP消息的结构

## （Request） 请求

Request（是个struct），代表了客户端发送的请求消息

重要字段：

+ URL
+ Header
+ Body
+ Form、PostForm、MultipartForm
也可以通过Request的方法访问请求中的Cookie、URL、User Agent等消息
Request既可以代表发送到服务器的请求，又可以代表客户端发出的请求

### 请求中的URL

Request 的URL字段就代表了请求行（请求信息第一行）里面的部分内容

URL 字段是指向url.URL类型的一个指针，url.URL是一个struct

```golang
type URL struct {
    Scheme   string
    Opaque   string
    User     *Userinfo
    Host     string
    Path     string
    RawQuery string
    Fragment string
}
```

### URL的通用形式

通用格式是：`scheme://[userinfo@]host/path[?query][#fragment]`

不以斜杆开头的URL会被解释成：
 `scheme:opaque[?query][#fragment]`

#### URL Query

RawQuery 会提供实际查询的字符串。

+ 例如： `http://www.example.com/post?id=123&thread_id=456`
+ 它的 RawQuery 的值就是 id=123&thread_id=456
+ r.URL.Query()，会提供查询字符串对应的 `map[string][]string`

例子：

```golang
url := r.URL
query := url.Query() //map[string][]string
id := query["id"] // []string{"123"}
threadID := query.Get("thread_id") //"456"
```

+ 还有一个简便方法可以得到 Key-Value 对：通过 Request 的 Form 字段

#### URL Fragment

+ 如果从浏览器发出的请求，那么你无法提取出 Fragment 字段的值
  + 浏览器在发送请求时会把 fragment 部分去掉
+ 但不是所有的请求都是从浏览器发出的（例如从 HTTP 客户端包）。

### Request Header

+ 请求和响应(request response)的headers是通过Header类型来描述的，它是一个map，用来表述HTTP Header里KEY-VALUE对

+ Header map 的key是string类型，value是[]string
+ 设置key的时候会创建一个空的[]string作为value，value里面第一个元素就是新header的值
+ 为指定的key添加一个新的header值，执行apped操作即可

### Request Header 例子

```golang
r.Header
 返回 map
r.Header["Accept-Encoding"]
 返回：【gzip，deflate】([]string 类型)
r.Header.Get["Accept-Encoding"]
 返回： gzip，deflate (string 类型)
```

### Request Body

请求和响应的bodies都是使用Body字段来表示的

Body是一个io.ReadCloser接口

+ 一个Reader接口
+ 一个Closer接口

Reader接口定义了一个Read方法：

+ 参数： []byte
+ 返回： byte的数量、可选的错误

Closer接口定义了一个Close方法：

+ 没有参数，返回可选的错误

## 示例代码

```golang
func main3() {
 http.HandleFunc("/header", func(rw http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(rw, r.Header)
  fmt.Fprintln(rw, r.Header["Accept-Encoding"])
  fmt.Fprintln(rw, r.Header.Get("Accept-Encoding"))
 })

 http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
  length := r.ContentLength
  body := make([]byte, length)
  r.Body.Read(body)
  fmt.Fprintln(rw, string(body))

 })

 http.ListenAndServe(":8080", nil)

}
```

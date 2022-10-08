# Form 表单

## 大纲

+ 通过表单发送请求
+ Form 字段
+ PostForm 字段
+ MultipartForm 字段
+ FormValue & PostFormValue 方法
+ 文件（Files）
+ POST JSON

## 来自表单的Post请求

```html
<form action="/process?" method="post" enctype="multipart/form-data">
        <input type="text" name="first_name"/>
        <input type="text" name="last_name"/>
        <input type="submit" />
</form>
```

+ 这个 HTML 表单里面的数据会以 name-value 对的形式，通过 POST 请求发送出去。
+ 它的数据内容会放在 POST 请求的 Body 里面

## 表单 Post 请求的数据格式

通过 POST 发送的 name-value 数据对的格式可以通过表单的 Content Type 来指定，也就是 enctype 属性。

### 表单的 enctype 属性

+ 默认值是：application/x-www-form-urlencoded
+ 浏览器被要求至少要支持： application/x-www-form-urlencoded 、multipart/form-data
+ HTML 5 的话，还需要支持 text/plain

+ 如果 enctype 是 application/x-www-form-urlencoded，那么浏览器会将表单数据编码到查询字符串里面。例如：
  + first_name=sau%20sheong&last_name=chang
+ 如果 enctype 是 multipart/form-data，那么
  + 每一个 name-value 对都会被转换为一个MIME消息部分
  + 每一个部分都有自己的 Content Type 和 Content Disposition

### 选择条件

+ 简单文本：表单 URL 编码
+ 大量数据，例如上传文件：multipart-MIME
  + 甚至可以把二进制数据通过选择 Base64 编码，来当作文本进行发送

## 表单的 GET

+ 通过表单的 method 属性，可以设置 POST 还是 GET

```html
<form action="/process?" method="get">
        <input type="text" name="first_name"/>
        <input type="text" name="last_name"/>
        <input type="submit" />
</form>
```

+ GET 请求没有 Body，所有的数据都通过 URL 的 name-value 对来发送

## Form字段

+ Request 上的函数允许我们从 URL 或/和 Body 中提取数据，通过这些字段：
  + Form
  + PostForm
  + MultipartForm
+ Form 里面的数据是 key-value 对。
+ 通常的做法是：
  + 先调用 ParseForm 或 ParseMultipartForm 来解析 Request
  + 然后相应的访问 Form、PostForm 或 MultipartForm 字段

## PostForm 字段

+ 前例中，如果只想得到 first_name 这个 Key 的 Value，可使用 r.Form[“first_name”]，它返回含有一个元素的 slice：[“Dave”]
+ 如果表单和 URL 里有同样的 Key，那么它们都会放在一个 slice 里：表单里的值靠前，URL 的值靠后
+ 如果只想要表单的 key-value 对，不要 URL 的，可以使用 PostForm 字段。
+ PostForm 只支持 application/x-www-form-urlencoded
+ 想要得到 multipart key-value 对，必须使用 MultipartForm 字段

## MultipartForm 字段

+ 想要使用 MultipartForm 这个字段的话，首先需要调用 `ParseMultipartForm` 这个方法
  + 该方法会在必要时调用 ParseForm 方法
  + 参数是需要读取数据的长度
+ MultipartForm 只包含表单的 key-value 对
+ 返回类型是一个 struct 而不是 map。这个 struct 里有两个 map：
  + key 是 string，value 是 []string
  + 空的（key 是 string，value 是文件）

## FormValue 和 PostFormValue 方法

+ FormValue 方法会返回 Form 字段中指定 key 对应的第一个 value
  + 无需调用 ParseForm 或 ParseMultipartForm
+ PostFormValue 方法也一样，但只能读取 PostForm
+ FormValue 和 PostFormValue 都会调用 ParseMultipartForm 方法
+ 但如果表单的 enctype 设为 multipart/form-data，那么即使你调用ParseMultipartForm 方法，也无法通过 FormValue 获得想要的值。

## 代码示例

```golang
func main() {
 server := http.Server{
  Addr: "localhost:9000",
 }

 http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
  // r.ParseForm()
  // fmt.Fprintln(w, r.Form)
  // fmt.Fprintln(w, r.PostForm)
  r.ParseMultipartForm(1024)
  fmt.Fprintln(w, r.MultipartForm)

 })

 server.ListenAndServe()
}
```

## Forms - MultipartReader()

### 读取Form的值

+ Form
+ PostForm
+ FormValue()
+ PostFormValue()
+ FormFile()
+ MultipartReader()

### MultipartReader()

+ `func (r *Request) MultipartReader() (*multipart.Reader, error)`
+ 如果是 multipart/form-data 或 multipart 混合的 POST 请求：
  + MultipartReader 返回一个 MIME multipart reader
  + 否则返回 nil 和一个错误
+ 可以使用该函数代替 ParseMultipartForm 来把请求的 body 作为 stream 进行处理
  + 不是把表单作为一个对象来处理的，不是一次性获得整个 map
  + 逐个检查来自表单的值，然后每次处理一个

## 上传文件

+ multipart/form-data 最常见的应用场景就是上传文件
  + 首先调用 ParseMultipartForm 方法
  + 从 File 字段获得 FileHeader，调用其 Open 方法来获得文件
  + 可以使用 ioutil.ReadAll 函数把文件内容读取到 byte 切片里

### 代码示例

```html
    <form action="http://localhost:9000/process?hello=world&thread=123" method="post" enctype="multipart/form-data">
        <input type="text" name="hello" value="hhhhha">
        <input type="text" name="post" value="456">
        <input type="file" name="uploaded">
        <input type="submit">
    </form>
```

```golang
package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
 r.ParseMultipartForm(1024)

 fileHeader := r.MultipartForm.File["uploaded"][0]
 file, err := fileHeader.Open()
 if err == nil {
  data, err := ioutil.ReadAll(file)
  if err == nil {
   fmt.Fprintln(w, string(data))
  }
 }
}

func main() {
 server := http.Server{
  Addr: "localhost:9000",
 }
 http.HandleFunc("/process", process)
 server.ListenAndServe()
}
```

### FormFile 方法

+ 上传文件还有一个简便方法：FormFile（例子）
  + 无需调用 ParseMultipartForm 方法
  + 返回指定 key 对应的第一个 value
  + 同时返回 File 和 FileHeader，以及错误信息
  + 如果只上传一个文件，那么这种方式会快一些

### 代码示例

```go
// 将上面例子中这部分替换
// r.ParseMultipartForm(1024)

// fileHeader := r.MultipartForm.File["uploaded"][0]
// file, err := fileHeader.Open()

file, _, err := r.FormFile("uploaded") //返回上传的第一个文件
```

## POST 请求 - JSON Body

+ 不是所有的 POST 请求都来自 Form
+ 客户端框架（例如 Angular 等）会以不同的方式对 POST 请求编码：
  + jQuery 通常使用 application/x-www-form-urlencoded
  + Angular 是 application/json
+ ParseForm 方法无法处理 application/json

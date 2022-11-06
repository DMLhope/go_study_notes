package main

import "text/template"

func main9() {
	//此函数的实现基本就是先将文件内容转换成字符串并将文件路径解析获得文件名不包含路径用来建一个新的模板，然后调用Parse方法来解析模板里的字符串
	// t, _ := template.ParseFiles("tmpl.html")
	//下面两句等价上面这句
	// t := template.New("tmpl.html")
	// t, _ = t.ParseFiles("tmpl.html")
	t, _ := template.ParseGlob("*.html")
}

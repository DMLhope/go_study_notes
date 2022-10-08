# 用Go实现一个简单的命令行工具

## 用os.Args获取命令行参数

+ os包提供了用于处理操作系统相关内容的函数/值
  + 独立于平台的方式
+ os.Args变量
  + 获取命令行参数
  + 他是个string slice（string类型的切片）
  + 第一个值是命令本身
+ strings.Join函数
  + strings包的Jion函数，能将string的集合组合成一个字符串

```golang
package main

import (
 "fmt"
 "os"
)

func main() {
 var s, sep string
 // os.Args 用这个获取命令行参数

 for i := 1; i < len(os.Args); i++ {
  s += sep + os.Args[i]
  sep = " "
 }

 fmt.Println(s)
}
```

+ 此处`i`从1开始是因为0为命令本身

## 代码优化

```golang
// 变量申明优化
var s, sep string
//可以优化成
s,sep := "", ""

// 循环换成range写法
for _, arg := range os.Args[1:] {
 s += sep + arg
 sep = " "
}

//不用循环的写法
fmt.Println(strings.Join(os.Args[1:], " "))

```

## 用户输入

+ bufio.NewReader()读取输入到缓存
+ os.Stdin读取命令行输入
+ reader.ReadString('\n')通过\n来判断输入结束

```golang
func main() {
 fmt.Println("Please input your words !!!")
 reader := bufio.NewReader(os.Stdin)
 text, _ := reader.ReadString('\n')
 fmt.Printf("Your words are : %s", text)
}
```

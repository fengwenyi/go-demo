# study-go

学习 Go

## Hello Go

程序代码：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}
```

控制台打印：
```
Hello Go!
```

## 基础

- [x] [00 hello go](00-hello-go)
- [x] [01 变量](01-varible)
- [x] [02 常量](02-constant)
- [x] [03 常量](03-func)
- [x] [04 并发编程](04-goroutine)
- [x] [05 switch](05-switch)
- [x] [06 goto](06-goto)
- [x] [07 数组](07-array)
- [x] [08 切片](08-slice)
- [x] [09 结构体](09-struct)
- [x] [10 指针](10-pointer)
- [x] [11 序列化](11-serialization)

## 数据类型

### 整型

### byte 与 rune  ???

byte与rune类型有一个共性，即：它们都属于别名类型。byte是uint8的别名类型，而rune则是int32的别名类型。

byte类型的值需用8个比特位表示，其表示法与uint8类型无异。因此我们就不再这里赘述了。我们下面重点说说rune类型。

一个rune类型的值即可表示一个Unicode字符。Unicode是一个可以表示世界范围内的绝大部分字符的编码规范。关于它的详细信息，大家可以参看其官网（http://unicode.org/）上的文档，或在Google上搜索。用于代表Unicode字符的编码值也被称为Unicode代码点。一个Unicode代码点通常由“U+”和一个以十六进制表示法表示的整数表示。例如，英文字母“A”的Unicode代码点为“U+0041”。
    rune类型的值需要由单引号“'”包裹。例如，'A'或'郝'。这种表示方法一目了然。不过，我们还可以用另外几种形式表示rune类型值。请看下表。  

大家需要根据实际情况选用上述表示法。在一般情况下，第一种表示法更为通用。因为它是最直观的。不过，在以其他几种方法表示的内容出现在屏幕上的时候，大家也要明白其含义。
  
另外，在rune类型值的表示中支持几种特殊的字符序列，即：转义符。它们由“\”和一个单个英文字符组成。如下表所示。


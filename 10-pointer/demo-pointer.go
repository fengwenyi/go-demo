package main

import "fmt"

func main() {
	//test1()
	test2()
}

func test1() {
	var name string = "张三"

	var namePointer *string = &name

	var age int = 18

	var agePointer *int = &age

	fmt.Println(name)
	fmt.Println(namePointer)
	fmt.Println(age)
	fmt.Println(agePointer)
}

// 指向的是内存地址
func test2() {
	name := "张三"
	namePointer := &name
	fmt.Println(name, namePointer, *namePointer)

	name = "李四"
	fmt.Println(name, namePointer, *namePointer)
	/*
		张三 0xc000010230 张三
		李四 0xc000010230 李四
	*/
}

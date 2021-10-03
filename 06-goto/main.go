package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			goto breakTag
		}
		fmt.Println("%v \n", i)
	}
breakTag:
	fmt.Println("结束遍历")
}

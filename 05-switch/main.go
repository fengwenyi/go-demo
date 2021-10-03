package main

import "fmt"

func main() {

	test1()
	test2()
	test3()

}

func test1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	}
}

func test2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("n为奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("n为偶数")
	}
}

func test3() {
	score := 30
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score > 80 && score < 90:
		fmt.Println("良好")
	case score > 90:
		fmt.Println("优秀")
	}
}

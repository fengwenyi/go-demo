package main

import "fmt"

func main() {
	//testInit()
	//arrayIterator()
	//manyArray()
	testModifyArrayValue()
}

func testInit() {
	init1()

	fmt.Println("-----------------------")

	init2()

	fmt.Println("-----------------------")

	init3()
}

func init1() {
	var arr1 [3]int
	var arr2 = [3]int{1, 2}
	var arr3 = [3]string{"北京", "上海", "深圳"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// 数组长度自动推导
func init2() {
	var arr1 [3]int
	var arr2 = [...]int{1, 2}
	var arr3 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("type of arr2:%T\n", arr2)
	fmt.Println(arr3)
	fmt.Printf("type of arr3:%T\n", arr3)
}

func init3() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)
}

func arrayIterator() {
	var array = [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	fmt.Println("-----------------------")
	for i, v := range array {
		fmt.Println(i, v)
	}
}

func manyArray() {
	cities := [3][2]string{
		{"四川", "成都"},
		{"广东", "广州"},
		{"湖南", "长沙"},
	}
	for _, v1 := range cities {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	fmt.Println("---------------")
	cities2 := [...][2]string{
		{"四川", "成都"},
		{"广东", "广州"},
		{"湖南", "长沙"},
	} // 第一层可以自动推断
	fmt.Println(cities2)
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyManyArray(x [3][2]int) {
	x[2][0] = 100
}

func testModifyArrayValue() {
	a := [3]int{10, 20, 30}
	modifyArray(a)
	fmt.Println(a)

	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyManyArray(b)
	fmt.Println(b)
}

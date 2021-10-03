package main

import "fmt"

func main() {
	//initSlice()
	//arraySlice()
	//arraySliceSlice()
	//completeSlice()
	//testMake()
	//sliceEmpty()
	//sliceEqual()
	//sliceDelete()
	//sliceCopy()
	//sliceAppend()
	sliceIterator()
}

func initSlice() {
	var a []string
	var b = []int{} // 声明切片并初始化
	var c = []bool{false, true}
	var d = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
	// fmt.Println(c == d) // 切片是引用类型，不支持直接比较，只能和nil比较
}

// 数组切片
func arraySlice() {
	//          0  1  2  3  4
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // 2,3
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	// 容量=得到切片底层数组的容量

	s2 := a[1:4]
	fmt.Println(cap(s2))
}

// 数组切片的切片
func arraySliceSlice() {
	//          0  1  2  3  4
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // 2,3
	s2 := s[1:2]
	fmt.Println(s2) // [3]
	s3 := s[3:4]
	fmt.Println(s3) // [5]
	s4 := s[0:1]
	fmt.Println(s4) // [2]
	s5 := s[1:3]
	fmt.Println(s5) // [3 4]
	//s6 := s[1:10] // 最大不能超过4
	//fmt.Println(s6) // 报错
}

// 切片的完整表达式
func completeSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3:5]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// s:[2 3] len(s):2 cap(s):4
}

// make()
func testMake() {
	s := make([]int, 2, 10)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

// 切片判空
func sliceEmpty() {
	var s1 []int
	s2 := []int{}
	fmt.Println(len(s1) == 0) // true
	fmt.Println(len(s2) == 0) // true
}

// 切片赋值
func sliceEqual() {
	s := []int{0, 0, 0}
	s2 := s
	s2[0] = 100
	fmt.Println(s)  // [100 0 0]
	fmt.Println(s2) // [100 0 0]
}

// 删除元素
func sliceDelete() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	// 删除3号元素 4
	s = append(s[:3], s[4:]...)
	fmt.Println(s) // 1，2，3，5
}

// 切片拷贝
func sliceCopy() {
	s := []int{1, 2, 3, 4, 5}
	t := make([]int, 5, 10)
	copy(t, s)
	fmt.Println(s)
	fmt.Println(t)
	s[0] = 1000
	fmt.Println(s)
	fmt.Println(t)
}

// 切片追加
func sliceAppend() {
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	s2 := []int{5, 6, 7}
	s = append(s, s2...)
	fmt.Println(s)
}

// 切片遍历
func sliceIterator() {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	fmt.Println("------------------")
	for i, v := range s {
		fmt.Println(i, v)
	}
}

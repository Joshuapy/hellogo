// 切片练习

package main

import "fmt"

func declare_slice() {
	//1 var 声明
	var slice1 []int
	isnil := slice1 == nil
	var slice2 = []int{1, 2}
	// slice1 = append(slice1, 1)
	fmt.Printf("slice1:%v, isnil: %v\n", slice1, isnil)
	fmt.Println("slice2:", slice2)
	// slice3 := []int  错误
	slice4 := []int{}
	slice5 := []int{1, 2, 3}
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

	// 2 from array
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice6 := arr1[:5]
	fmt.Println("slice6:", slice6)

	// 3 make()
	slice7 := make([]int, 3, 6)
	fmt.Println("slice7:", slice7)
	fmt.Print(len(slice7), cap(slice7))
	var slice8 = make([]int, 5)
	fmt.Println("slice8:", slice8)
}

func append_elem() {
	slice1 := []int{}
	fmt.Printf("%p\n", &slice1)
	slice1 = append(slice1, 1)
	fmt.Printf("%p\n", &slice1)
	fmt.Printf("len: %d, cap: %d, %v\n", len(slice1), cap(slice1), slice1)
	//fmt.Printf("len: %d, cap: %d, %v\n", len(slice2), cap(slice2), slice2)
}

func slice_cap() {
	s := make([]int, 2, 2)
	s = append(s, 1)
	fmt.Printf("%p\n", &s[0])
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", &s[0])
	fmt.Println(len(s), cap(s))
}

func slice_data() {
	arr1 := [...]int{1, 2, 3, 10: 0}
	slice1 := arr1[:2]
	slice2 := arr1[:5]
	slice1[0] = 10
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func append_slice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := append(slice1, slice2...)
	fmt.Printf("%v %d %d", slice3, len(slice3), cap(slice3))
}

func slice_copy() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 5)
	fmt.Printf("slice1: len: %d, cap: %d, %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2: len: %d, cap: %d, %v\n", len(slice2), cap(slice2), slice2)
	copy(slice2, slice1)
	fmt.Printf("slice1: len: %d, cap: %d, %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2: len: %d, cap: %d, %v\n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 5)
	fmt.Printf("slice2: len: %d, cap: %d, %v\n", len(slice2), cap(slice2), slice2)

}

// 遍历
func range_slice() {
	aslice := make([]int, 6)  // 初始零值
	aslice[0] = 1
	aslice = append(aslice, 11)
	// aslice := []int{1, 1, 1, 2, 3}
	// for i := range aslice {  // 注意这里的变量： 一个的时候是索引，两个的时候是索引+元素
	for i, v := range aslice {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
}

func main() {
	declare_slice()
	//append_elem()
	//slice_cap()
	// slice_data()
	// append_slice()
	// slice_copy()
	range_slice()
}

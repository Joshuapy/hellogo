// 数组练习

package main

import "fmt"

func declare_array() {
	arr1 := [...]int{1, 2, 3}
	fmt.Print(arr1)
}

func access_element() {
	var arr1 = [3]int{}
	arr1[0] = 3
	fmt.Print(arr1)
}

func main() {
	//declare_array()
	access_element()
}

package main

import "fmt"

func test1() {
	var a int
	fmt.Printf("a:%d\n", a)
	var b *int
	// b = &a
	*b = 20
	fmt.Printf("a:%d\n", a)
}


// 空指针
func empty_ptr() {
	var ptr *int
	fmt.Printf("ptr is: %#v\n", ptr) // ptr is <nil>
	// fmt.Printf("*ptr is: %#v\n", *ptr) // panic, 空指针没法取值
}

func main() {
	// test1()
	empty_ptr()
}

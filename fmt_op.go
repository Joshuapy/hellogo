// fmt内置包练习
// 2023.10.18

package main

import (
	"fmt"
	"os"
)

// print系列, 直接输出到终端
func print() {
	fmt.Print("直接输出", 1, true, 2.1, "\n")
	fmt.Println("换行输出")
	fmt.Printf("格式化输出: %s\n", "填充值")
}

// io.Writer输出
func fprint() {
	fileObj, err := os.OpenFile("./fmt_op_test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error")
	}
	fmt.Fprint(fileObj, "test for Fprint\n")
	fmt.Fprintf(fileObj, "test for %s\n", "Fprintf")
	fmt.Fprintln(fileObj, "test for Fprintln")

	fmt.Println("fprint done")
}

// sprint 返回字符串
func sprint() {
	s1 := fmt.Sprint("1", "2") // 无缝拼接后返回一个字符串
	s2 := fmt.Sprintf("1:%s, 2:%s\n", "1", "2")
	s3 := fmt.Sprintln("带换行xxx")
	fmt.Printf("s1: %q\n", s1)
	fmt.Printf("s2: %q\n", s2)
	fmt.Printf("s3: %q\n", s3)
}

func format_str() {
	name := "joker"
	age := 18
	tall := 8.88
	isOk := true
	p := &name
	type T struct {
		name string
		age  int
	}
	t := T{"imli", 18}

	fmt.Printf("name: %s, age: %d, tall: %f, isMan: %t, t: %v\n", name, age, tall, isOk, t)
	fmt.Printf("name: %q, age: %b, 'tall: %10.2f', isMan: %v, t: %#v\n", name, age, tall, isOk, t)
	fmt.Printf("name: %v, age: %v, 'tall: %v', isMan: %v, t: %+v\n", name, age, tall, isOk, t)
	fmt.Printf("name: %#v, age: %#v, 'tall: %#v', isMan: %#v\n", name, age, tall, isOk)
	fmt.Printf("pointer: %v, %p, %#p\n", p, p, p)
}

func main() {
	// print()
	// sprint()
	// fprint()
	format_str()
}

// 引包练习
// 2022/4/29

package main

import "fmt"
import foo "hellogo/pkgfoo" // 子包路径，而非包名; hellgo 是模块前缀，来自于go.mod的模块名定义

func main() {
	fmt.Printf("this is in main package,and import varialbe from package foo: %v\n", foo.Foo)
}

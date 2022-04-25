// Description 练习变量明和赋值
// Date 2021-11-17
// Author: joshua

package main

import (
	"fmt"
)

// 格式不一样而已 看习惯
var (
	nameN int // 需在函数内赋值
	nameM string
)

// 1变量申明
var abc int32 // 有默认值空值

// 2变量声明 + 赋值
var name string = "my name is wanghua!"
var count int = 100

// 3变量声明 + 省类型赋值
var age_str = "32"
var age = 32

// --------------------------------------------------
// 1
var name1, name2, name3 int // 标准变量声明， 只能在函数体内赋值
// 2
var age1, age2, age3 int = 1, 2, 3 // 声明+赋值
// 3
var age11, age22, age33 = 1, 2, 3 // 自动判断类型, 省去类型申明

func print_some() {
	name1, name2, name3 = 1, 2, 3 // 对应 1 申明的赋值

	// 4
	count1, count2, count3 := 11, 22, 33 // 只能局部变量使用该方式

	fmt.Println(name1, name2, name3)
	fmt.Println(age1, age2, age3)
	fmt.Println(age11, age22, age33)
	fmt.Println(count1, count2, count3)

	// ----------------
	// 支持交换变量
	a, b := 1, 2
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}

func main() {
    var u8 uint8
    u8 = 300
	fmt.Println("abc:", abc) // 默认值空值0

	abc = 111

	// 4变量局部申明 + 赋值
	number := "1111"
	number = "xxxx" //局部内可覆盖

	fmt.Println("abc:", abc)
	fmt.Println("number:", number)
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("age_str:", age_str)
	fmt.Println("count:", count)

	fmt.Println("nameN:", nameN)
	fmt.Println("nameM:", nameM)

	print_some()
}

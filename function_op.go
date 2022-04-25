// 函数练习
// 2021-4-25

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// 可变参数调用
	fmt.Println("可变参数:")
	sum_str := get_sum("xx", 1, 2, 3)
	fmt.Printf("%s\n", sum_str)

	// b := [...]int{1, 2, 3} // 不支持数组
	sliceInt := []int{1, 2, 3}
	sum_str = get_sum("yy", sliceInt...) // 切片解包
	fmt.Printf("%s\n", sum_str)

	// 多返回值
	fmt.Println("mutl_returnval:")
	aa, bb, cc := multi_returnval(10, 5) // 具名返回值, 推荐多用这个形式, 简洁，可读性友好
	fmt.Printf("sum: %v, cha: %v, ji: %v\n", aa, bb, cc)
	aaa, bbb, ccc := multi_returnval2(10, 5)
	fmt.Printf("sum: %v, cha: %v, ji: %v\n", aaa, bbb, ccc)

	// error 返回
	fmt.Println("error判断:")
	res, err := num_sprt(-1)
	fmt.Printf("res: %.2f, err: %v\n", res, err)
	res, err = num_sprt(4)
	fmt.Printf("res: %.2f, err: %v\n", res, err)
	res, err = num_sprt2(-10)
	fmt.Printf("res: %.2f, err: %v\n", res, err)
	res, err = num_sprt2(9)
	fmt.Printf("res: %.2f, err: %v\n", res, err)

	// defer
	defer_basic()
	defer_tmpvarable()
	defer_closure_var()
	defer_closure_params()
}

// 函数基本定义
func basic_function(a, b int) int {
	c := a + b
	return c
}

//可变参数(切片)
func get_sum(a string, b ...int) string {
	var x int
	for _, i := range b {
		x += i
	}
	msg := fmt.Sprintf("get_sum: a: %s: %d", a, x)
	return msg
}

// 多返回值-具名返回值
func multi_returnval(a, b int) (sum, cha, ji int) {
	sum = a + b
	cha = a - b
	ji = a * b
	return
}

// 多返回值-非具名返回值
func multi_returnval2(a, b int) (int, int, int) {
	sum := a + b
	cha := a - b
	ji := a * b
	return sum, cha, ji
}

// 判断，返回error
func num_sprt(n float64) (res float64, err error) { // 具名返回值
	if n < 0 {
		err = errors.New("too litle!")
		return
	}
	res = math.Sqrt(n) // 平方根
	return res, nil
}
func num_sprt2(n float64) (float64, error) { // 非具名返回值
	if n < 0 {
		return 0.0, errors.New("too litle!")
	}
	res := math.Sqrt(n)
	return res, nil
}

// ================= defer 延迟执行  ================
func defer_basic() {
	var arr1 [5]int
	for i, _ := range arr1 {
		defer fmt.Println(i) // 4 3 2 1 0  正序注册，倒序执行
	}
}

// 临时变量注册的坑
type SomeConn struct {
	name string
}

func (c *SomeConn) close() {
	fmt.Printf("name: %v closed\n", c.name)
}
func defer_tmpvarable() {
	conn := []SomeConn{{name: "a"}, {name: "b"}, {name: "c"}}
	for _, c := range conn {
		tmp_conn := c          // 修复，类似range的  取地址运算，注册的临时变量导致。
		defer tmp_conn.close() // c c c
	}
}

// defer 闭包变量引用
func defer_closure_var() {
	var slice1 = make([]int, 5)
	for i := range slice1 {
		defer func() { fmt.Println(i) }() // 闭包变量属于引用 output: 44444
	}
}

//defer 闭包参数是复制
func defer_closure_params() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Printf("defer: x=%v, y=%v\n", i, y) // 外部变量是引用
	}(x) // 值传递(复制了一遍)
	x += 10
	y += 100
	fmt.Printf("x=%v, y=%v\n", x, y) // x=20, y=120
}

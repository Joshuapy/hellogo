// 函数

package main

import (
	"errors"
	"fmt"
	"math"
)

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
func num_sprt(n float64) (res float64, err error) {
	if n < 0 {
		err = errors.New("too litle!")
		return 0.0, err
	}
	res = math.Sqrt(n)
	return res, nil
}
func num_sprt2(n float64) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("too litle!")
	}
	res := math.Sqrt(n)
	return res, nil
}

func main() {
	// res := get_sum("xx", 1, 2, 3)
	// fmt.Printf("%s\n", res)

	// // b := [...]int{1, 2, 3} // 不支持数组
	// b := []int{1, 2, 3}       // 不支持数组
	// res = get_sum("yy", b...) // 解包
	// fmt.Printf("%s\n", res)

	// fmt.Println("mutl_returnval:")
	// aa, bb, cc := multi_returnval(10, 5)
	// fmt.Printf("sum: %v, cha: %v, ji: %v\n", aa, bb, cc)
	// aaa, bbb, ccc := multi_returnval2(10, 5)
	// fmt.Printf("sum: %v, cha: %v, ji: %v\n", aaa, bbb, ccc)

	// error 返回
	res, err := num_sprt2(10)
	fmt.Printf("res: %.2f, err: %v\n", res, err)
}

// 错误处理练习
// 2022/4/26

package main

import (
	"errors"
	"fmt"
)

func main() {
	what_is_error()
	handle_error()
	// panic_occurred()

	i, err := defer_recover()
	fmt.Printf("result of defer_recover: %d, err: %v\n", i, err)

	fmt.Printf("Main end!\n")
}

// 错误对象构造
func what_is_error() {
	err := errors.New("this is a normal error message.")
	err2 := fmt.Errorf("this is a fmt error: %s", "fmt.Errorf")
	fmt.Printf("Error: %s\n", err)
	fmt.Printf("Error2: %s\n", err2.Error())
	fmt.Printf("type of err: %#T\n", err)   // *errors.errorString
	fmt.Printf("type of err2: %#T\n", err2) // *errors.errorString
}

func return_err(num int) (string, error) {
	var msg string
	if num < 0 {
		return msg, errors.New("too small!")
	}
	msg = "ok"
	return msg, nil
}

// 处理错误范式
func handle_error() {
	n := -10
	value, err := return_err(n)
	if err != nil { // 只判断错误对象需要想好如何中止。
		fmt.Printf("Error with parameter: %d: %s\n", n, err.Error())
	}
	fmt.Printf("success: %s\n", value) // 依然会执行 value是空值

	// 推荐使用如下if-else 复式判断
	m := -10
	if value, err := return_err(m); err != nil {
		fmt.Printf("Error with parameter: %d: %s\n", m, err.Error())
	} else {
		fmt.Printf("success: %s\n", value)
	}

}

// 触发panic: 1 主动触发； 2 运行时异常
func panic_occurred() {
	fmt.Println("Panic Starting!")
	panic("Some Server error occured!") //普通错误 [panic: Some Server error occured!]  任意类型参数;通常是一个字符串.
	// a := []int{1,2}
	// b := a[3] // 下标越界 运行时错误 [panic: runtime error: index out of range [3] with length 2]
	// fmt.Printf("b: %v\n", b)
	fmt.Println("End!\n") // 这个不会执行
}

// panic-defer-revoer 机制
// 异常捕获，目的是中止panicking, 并使整个程序继续运行
func defer_recover() (res int, err error) {
	defer func() { // 设置异常拦截站，防止继续panicking
		if ferr := recover(); ferr != nil {
			fmt.Printf("function: panic_occurred error: %s, type of ferr: %T\n", ferr, ferr)
			// err = ferr // ferr 不是一个标准的error类型,而是一个接口类型
			err = errors.New(fmt.Sprint(ferr)) // need to cover interface{} to string
			return
		}
	}()
	panic_occurred()
	fmt.Println("Still running!") // 也不会执行？
	return 1, nil
}

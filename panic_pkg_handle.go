// 一个完整的包内异常处理范例
// 2022/4/26
// 最佳实践:
/* 
    1. 包内部总是从panic中recover; 不允许显示的超出包范围的panic(); 包的顶级函数都应该有defer-recover机制
    2. 包外的调用者只会收到错误值，而不是panic
*/

package main

import (
    "fmt"
    "strings"
    "strconv"
)


func main() {
    input := "xx yy"
    // input := "1 2 3"
    out, err := Parse(input)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Out: %v\n", out)
    }
}

type ParseError struct {
    Index int
    Word string
    Err error
}

// 方法
func (e *ParseError) String() string {
    return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// 包外可见
func Parse(input string) (numbers []int, err error) {
    defer func(){
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok {
                err = fmt.Errorf("pkg: %v", r)
            }
        }
    }()

    fields := strings.Fields(input)
    numbers = fields2numbers(fields)
    return
}

// 内部函数，可以做显式panic操作
func fields2numbers(fields []string) (numbers []int) {
    if len(fields) == 0 {
        panic("no words to parse!")
    }

    for idx, field := range fields {
        num, err := strconv.Atoi(field)
        if err != nil {
            panic(&ParseError{idx, field, err})
        }
        numbers = append(numbers, num)
    }
    return
}


// 练习runtime包的使用
// 2023.10.11

package main

import (
	"fmt"
	"runtime"
	"time"
)

// 主动让出分片，等待再次分配执行分片
// xxx
// yyy
func Gosched() {
	go func(s string) {
		for i := 0; i < 100; i++ {
			fmt.Println(s, i)
		}
	}("world")

	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
		runtime.Gosched()
	}
}

// Goexit()  主动退出协程的执行，已经注册的defer会依次执行。
func GoExist() {
	go func() {
		defer fmt.Println("A.defer") // 2
		func() {
			defer fmt.Println("B.defer") // 1
			runtime.Goexit()
			defer fmt.Println("C.defer") //已经退出协程,defer未注册
			fmt.Print("B")
		}()
		fmt.Println("A")
	}()

	for {
	}

}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("a", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("b", i)
	}
}
func GoMaxprocs() {
	// 设置一个核心执行
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

func main() {
	// GoExist()
	// Gosched()
	GoMaxprocs()
}

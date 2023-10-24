// 练习channel基本操作
// 2023.10.12

package main

import (
	"fmt"
	"time"
)

// ======== 无缓冲通道================
func recv(c chan int) {
	ret := <-c
	fmt.Println("接受成功", ret)
}

func sender() {
	ch := make(chan int) // 无缓冲通道
	go recv(ch)
	ch <- 10 // (无缓冲通道)发送之前必须有goroutine做好接受监听
	fmt.Println("发送成功")
	// go recv(ch) // 放在这里会panic
}

// ============有缓冲通道=============
func sender_buffer() {
	n := 5
	ch := make(chan int, 1)
	ch <- n
	fmt.Println("send success")
	go recv(ch)
	time.Sleep(2 * time.Second)
}

// ==========循环取值===========
func recv_loop1() {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1) // 这里必须显示关闭，否则接收端也会发生deadlock
	}()
	for i := range ch1 { // 读取完后会退出for循环
		fmt.Println(i)
	}
	fmt.Println(<-ch1) // 仍然可以读，不过返回的是零值
}
func recv_loop2() {
	ch1 := make(chan int)
	go func() {
		for i := 100; i < 110; i++ {
			ch1 <- i
		}
		close(ch1) // 这里必须显示关闭，否则接收端也会发生deadlock
	}()
	for {
		i, ok := <-ch1 // 关闭后的通道再取值会ok=false
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

// ============单向通道================
// send-only
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	// i := <-out  // send-only error
	close(out)
}

// 从in里读出来，处理后发送到out
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	// in <- 10 // receive-only error
	close(out)
}

// 只读
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func run() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func main() {
	// sender()
	// sender_buffer()
	recv_loop1()
	// recv_loop2()
	// run()
}

// 练习select
// 2023.10.17

package main

import (
	"fmt"
	"time"
)

func test3(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go test3(ch1)
	go test2(ch2)

	for i := 0; i < 100; i++ {
		select {
		case s1 := <-ch1:
			fmt.Println("s1=", s1)
		case s2 := <-ch2:
			fmt.Println("s2=", s2)
		}
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

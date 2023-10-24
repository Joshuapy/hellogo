package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("this is a test")
	log.Fatal("test log")
	fmt.Println("never shoud output")
}

package main

import (
	"fmt"
	"time"
)

func counter(kind string) {
	for i := 0; i < 5; i++ {
		fmt.Println(kind, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go counter("a")
	go counter("b")
	time.Sleep(time.Second * 10)
}

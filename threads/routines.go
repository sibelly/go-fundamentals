package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(msg string) {
	fmt.Println(msg + "_goroutine")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}

func main() {
	go hello("hello 1")
	go hello("hello 2")

	time.Sleep(1 * time.Second)
	fmt.Println("chamada normal")
}

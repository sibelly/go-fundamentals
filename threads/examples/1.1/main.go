package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Starting...")

	// Here go can work in a preemptive way (changing the context, since go1.14), killing this 'forever thread'
	go func() {
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Finishing...")
}

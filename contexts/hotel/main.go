package main

import (
	"context"
	"fmt"
	"time"
)

/*
context.Background() -> empty context, father of all
context.TODO() -> empty context that can be used some time in the future
context.WithCancel() -> returns a cancel() function that can be called when needed
context.WithTimeout() -> in some time will expire
context.WithDeadline() -> in some time plus more time will expire
context.WithValue("key", "value") -> can pass values from one side to another
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	// Another thread in background
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Time exceeded to book the room")
	case <-time.After(time.Second * 5):
		fmt.Println("Room successfuly booked")
	}
}

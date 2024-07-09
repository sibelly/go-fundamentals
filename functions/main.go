package main

import "fmt"

// Learning about tricky funcs
func main() {
	result := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}

	// Return the mem address of the inner func
	fmt.Println(result(1, 2, 3, 5))

	// Return the value of the inner func
	// () that calls the inner func indeed and return the value
	fmt.Println(result(1, 2, 3, 5)())
}

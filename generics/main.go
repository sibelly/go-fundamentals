package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ~ symbol does an approximation for the type
type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SumGeneric[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// comparable constraint
// The comparable interface may only be used as a type parameter constraint, not as the type of a variable
func Sum[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >
// https://pkg.go.dev/golang.org/x/exp/constraints
func Max[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

// ##########################################
type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

//##########################################

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3

	fmt.Println(SumGeneric(map[string]MyNumber{"a": x, "b": y, "c": z}))
	fmt.Println(SumGeneric(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))

	// concat
	fmt.Println(concat([]MyString{"a", "b", "c"}))
}

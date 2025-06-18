package main

import (
	"fmt"
	generic01 "hello_world/Generic01"
)

type Stack01[T comparable] = generic01.Stack[T]

func main() {
	// fmt.Println("Hello world welcome 123")
	// example01.MyExample01()

	// arrays01.Array01Func()

	var s1 Stack01[int]

	s1.Push(10)
	s1.Push(20)

	fmt.Println("generic ", s1)
	fmt.Println("contains 10", s1.Contains(10))

}

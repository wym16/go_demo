package main

import (
	"fmt"
	"unsafe"
)

const (
	A = "A"
	B = len(A)
	C = unsafe.Sizeof(A)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("area = %d", area)
	println()
	println(a, b, c)

	fmt.Println(A, B, C)
}

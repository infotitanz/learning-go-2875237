package main

import "fmt"

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println(*p)

	anotherInt := 7
	anotherPointer := &anotherInt
	fmt.Println("Another Value: ", *anotherPointer)

	*anotherPointer = *anotherPointer + 1
	fmt.Println("Another Pointer: ", *anotherPointer)
	fmt.Println("Another Value: ", anotherInt)

}

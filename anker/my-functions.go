package main

import "fmt"

func main() {
	sayHello()
	sum := addValues(1, 2)
	fmt.Println("Result of simple sum: ", sum)

	allSum, noOfValues := addAllValues(1, 2, 3, 4, 5, 9)
	fmt.Println("Result of Multi Sum: ", allSum)
	fmt.Println("Number of values provided: ", noOfValues)
}

func sayHello() {
	fmt.Println("Hello")
}

func addValues(a, b int) int {
	return a + b
}

func addAllValues(values ...int) (int, int) {
	result := 0
	for _, v := range values {
		result += v
	}
	return result, len(values)
}

package main

import (
	"fmt"
	"strconv"
)

func calculate(value1 string, value2 string) float64 {
	value1Float, err := strconv.ParseFloat(value1, 64)
	value2Float, err := strconv.ParseFloat(value2, 64)
	sum := 0.0
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		sum = value1Float + value2Float
	}

	return sum
}

func main() {
	fmt.Println(calculate("dani", "ill"))
}

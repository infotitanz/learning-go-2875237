package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"

	result, err := calculate(value1, value2, operation)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}

func calculate(input1, input2, operation string) (float64, error) {
	val1, err := convertInputToValue(input1)
	if err != nil {
		return 0, err
	}

	val2, err := convertInputToValue(input2)
	if err != nil {
		return 0, err
	}

	switch operation {
	case "+":
		return addValues(val1, val2), nil
	case "-":
		return subtractValues(val1, val2), nil
	case "*":
		return multiplyValues(val1, val2), nil
	case "/":
		if val2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return divideValues(val1, val2), nil
	default:
		return 0, fmt.Errorf("invalid operation: %s", operation)
	}
}

func convertInputToValue(input string) (float64, error) {
	floatInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}
	return floatInput, nil
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

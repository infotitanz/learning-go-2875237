package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	i1, i2, i3 := 25, 42, 63
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3 := 25.8, 42.7, 63.4
	floatSum := f1 + f2 + f3

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float Sum:", floatSum)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter age: ")
	input, _ := reader.ReadString('\n')
	age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age Next Year:", age+1)
	}
}

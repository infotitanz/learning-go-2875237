package main

import (
	"fmt"
	"sort"
)

func main() {

	var countriesVisited = []string{"Nigeria", "Ghana", "Dubai", "Ethiopia", "America", "Denmark", "Sweden", "Norway", "Germany", "Canada"}
	fmt.Println(countriesVisited)

	countriesVisited = append(countriesVisited, "France", "Spain")
	fmt.Println(countriesVisited)

	countriesVisited = countriesVisited[1:]
	fmt.Println(countriesVisited)

	countriesVisited = countriesVisited[:len(countriesVisited)-1]
	fmt.Println(countriesVisited)

	numbers := make([]int, 5)

	numbers[0] = 1
	numbers[1] = 5
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 2
	fmt.Println(numbers)

	numbers = append(numbers, 6)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}

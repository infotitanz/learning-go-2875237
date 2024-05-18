package main

import (
	"fmt"
	"math"
)

func main() {

	fruits := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
	// Call the convertToMap function with the fruits slice
	result := convertToMap(fruits)
	fmt.Println(result)

}

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)

	// get the length of the slice of strings and calculate the price of each item
	numberOfItems := len(items)
	itemPrice := 100.0 / float64(numberOfItems)
	itemPrice = math.Round(itemPrice*100) / 100

	// loop through the slice of strings
	for _, item := range items {
		result[item] = itemPrice
	}

	// Your code goes here
	return result
}

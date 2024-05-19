package main

import (
	"encoding/json"
)

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
		  {"name":"orange","price":1.50,"quantity":8},
		  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	for _, item := range result {
		println("Name: ", item.Name)
		println("Price: ", item.Price)
		println("Quantity: ", item.Quantity)
	}
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	// read the json string and return a slice of cartItem objects
	content := []byte(jsonString)
	err := json.Unmarshal(content, &cart)
	if err != nil {
		panic(err)
	}
	// Your code goes here.
	return cart
}

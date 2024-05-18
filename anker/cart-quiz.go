package main

import "fmt"

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}

	cart = append(cart, apples, oranges, bananas)

	// print cart total format to 2 decimal places
	fmt.Printf("%.6f", calculateTotal(cart))

}

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	total := 0.0
	for _, item := range cart {
		itemTotal := item.price * float64(item.quantity)
		total += itemTotal
	}
	return total
}

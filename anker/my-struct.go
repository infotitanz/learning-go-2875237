package main

import "fmt"

func main() {
	corolla := Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Price: 20000,
	}

	fmt.Println("Brand:", corolla.Brand)
	fmt.Println("Model:", corolla.Model)
	fmt.Println("Year:", corolla.Year)
	fmt.Println("Price:", corolla.Price)

	corolla.Price = 25000
	fmt.Println("Price:", corolla.Price)
}

type Car struct {
	Brand string
	Model string
	Year  int
	Price float64
}

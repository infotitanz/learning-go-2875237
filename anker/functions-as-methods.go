package main

import "fmt"

func main() {
	corolla := &Car{"Toyota", "Corolla", 2020, 20000, false}
	corolla.display()
	corolla.setPrice(22000)
	corolla.display()
}

type Car struct {
	Brand string
	Model string
	Year  int
	Price float64
	isSuv bool
}

func (c *Car) display() {
	fmt.Println("Brand:", c.Brand, "Model:", c.Model, "Year:", c.Year, "Price:", c.Price, "Is SUV:", c.isSuv)
}

func (c *Car) setPrice(price float64) {
	c.Price = price
}

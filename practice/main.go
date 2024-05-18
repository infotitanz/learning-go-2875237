package main

import "fmt"

func main() {
	aString := "Hello from Go!"
	anInt := 42
	fmt.Println(aString, "\nVariable Type:", fmt.Sprintf("%T", aString), "\nVariable Type:", fmt.Sprintf("%T", anInt))
}

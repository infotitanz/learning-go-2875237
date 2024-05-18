package main

import "fmt"

func main() {

	age := 15
	var result string

	if age >= 18 {
		result = "You are allowed to vote"
	} else {
		result = "You are not allowed to vote"
	}

	fmt.Println(result)
}

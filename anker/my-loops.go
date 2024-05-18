package main

func main() {

	colours := []string{"red", "green", "blue"}

	// for loop using range
	for _, colour := range colours {
		println(colour)
	}

	value := 1
	for value < 10 {
		println(value)
		value++
	}
}

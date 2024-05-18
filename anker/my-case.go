package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	dayOfWeek := rand.Intn(7) + 1

	var result string
	switch dayOfWeek {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	case 6:
		result = "Saturday"
	case 7:
		result = "Sunday"
	default:
		result = "Invalid day"
	}

	fmt.Println("Day of week: ", result)
}

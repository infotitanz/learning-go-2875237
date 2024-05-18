package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Today is", now.Weekday(), now.Day(), now.Month(), now.Year())
	fmt.Println(now.Format(time.ANSIC))

	parsedTime, err := time.Parse(time.ANSIC, "Fri May 17 15:10:42 2024")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(parsedTime)
	}

}

package main

import (
	"fmt"
	"sort"
)

func main() {

	provinces := make(map[string]string)

	provinces["AB"] = "Alberta"
	provinces["BC"] = "British Columbia"
	provinces["MB"] = "Manitoba"

	// Print the provinces
	fmt.Println(provinces)
	fmt.Print(provinces["AB"] + "\n" + provinces["BC"] + "\n" + provinces["MB"] + "\n")

	delete(provinces, "AB")
	fmt.Println(provinces)

	provinces["NB"] = "New Brunswick"
	provinces["NL"] = "Newfoundland and Labrador"
	provinces["NS"] = "Nova Scotia"
	provinces["ON"] = "Ontario"
	provinces["PE"] = "Prince Edward Island"
	provinces["QC"] = "Quebec"
	provinces["SK"] = "Saskatchewan"
	provinces["YT"] = "Yukon"
	provinces["NT"] = "Northwest Territories"
	provinces["NU"] = "Nunavut"
	provinces["AB"] = "Alberta"

	// loop through the map
	for key, value := range provinces {
		fmt.Println(key, value)
	}

	// loop through the map and display in alphabetical order
	keys := make([]string, 0, len(provinces))
	for key := range provinces {
		keys = append(keys, key)
	}

	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, provinces[key])
	}

}

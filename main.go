package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New delhi",
	}
	fmt.Println("origin map")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])

		delete(countryCapitalMap, "France")
		fmt.Println("delete France")

		for country := range countryCapitalMap {
			fmt.Println(country, "首都是", countryCapitalMap[country])
		}
	}
}

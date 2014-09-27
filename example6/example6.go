package main

import (
	"fmt"
)

//	Returns n value?! Not sure how I feel about this...
func location(city string) (string, string) {
	var region string
	var country string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, country = "California", "North America"

	case "New York", "NYC":
		region, country = "New York", "North America"

	default:
		region, country = "Unknown", "Unknown"
	}

	return region, country
}

func main() {
	region, country := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, country)
}

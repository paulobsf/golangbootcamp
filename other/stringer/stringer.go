package main

import "fmt"

type Dish struct {
	Name        string
	Ingredients []string
	Price       float32
	Rating      int
}

func (d *Dish) String() string {
	stars := func(n int) string {
		var stars string
		for i := 0; i < n; i++ {
			stars += "★"
		}
		return stars
	}(d.Rating)

	return fmt.Sprintf("[%s] %s (%s): £ %f", stars, d.Name, d.Ingredients, d.Price)
}

func main() {
	d := new(Dish)
	d.Name = "Piri Piri Chicken"
	d.Ingredients = []string{
		"Chicken",
		"Piri Piri Sauce",
	}
	d.Price = 8.50
	d.Rating = 4

	//	Output: [★★★★] Piri Piri Chicken ([Chicken Piri Piri Sauce]): £ 8.500000
	fmt.Printf("%s\n", d)
}

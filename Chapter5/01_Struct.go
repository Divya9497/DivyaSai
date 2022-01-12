package main

import "fmt"

func main() {
	type person struct {
		first    string
		last     string
		favflavs []string
	}
	type Tata struct {
		brand   []string
		doors   int
		mileage int
	}
	type Hyundai struct {
		brand []string
		doors int
		year  []int
	}
	p1 := person{
		first:    "Divya",
		last:     "Sai",
		favflavs: []string{"Butterscotch", "chocolate"},
	}
	p2 := person{
		first:    "Sai",
		last:     "xyz",
		favflavs: []string{"vanilla", "Black Current"},
	}
	t1 := Tata{
		brand:   []string{"Nexon", "Tiago"},
		doors:   4,
		mileage: 20,
	}
	h1 := Hyundai{
		brand: []string{"verna", "aura"},
		doors: 4,
		year:  []int{2009, 2011},
	}
	fmt.Println(h1)
	fmt.Printf("%d\n", h1.doors)
	fmt.Println(t1.brand)
	fmt.Println(t1.doors)
	fmt.Println(t1.mileage)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.favflavs)
	for i, v := range p1.favflavs {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	fmt.Println(p2.favflavs)
	for i, v := range p2.favflavs {
		fmt.Println(i, v)
	}
}

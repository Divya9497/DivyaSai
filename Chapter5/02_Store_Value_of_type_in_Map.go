package main

import "fmt"

func main() {
	type person struct {
		first    string
		last     string
		favflavs []string
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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for i, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		fmt.Println(i)
		for j, val := range v.favflavs {
			fmt.Println(j, val)
		}
	}

}

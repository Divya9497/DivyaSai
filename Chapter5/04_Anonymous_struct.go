package main

import "fmt"

func main() {

	p1 := struct {
		first    string
		last     string
		favflavs []string
	}{
		first:    "Divya",
		last:     "Sai",
		favflavs: []string{"Butterscotch", "chocolate"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.favflavs)
	for i, v := range p1.favflavs {
		fmt.Println(i, v)
	}
}

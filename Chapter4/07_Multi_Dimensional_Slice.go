package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(x)
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(y)
	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println("slice:", i)
		for j, v1 := range v {
			fmt.Println(j, v1)
		}
	}
}

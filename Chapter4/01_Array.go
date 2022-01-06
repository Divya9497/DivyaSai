package main

import "fmt"

func main() {
	x := [5]int{}
	x[0] = 23
	x[1] = 32
	x[2] = 45
	x[3] = 54
	x[4] = 65

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

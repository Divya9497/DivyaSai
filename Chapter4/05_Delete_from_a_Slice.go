package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := x[:3]
	fmt.Println(y)
	z := x[6:]
	fmt.Println(z)
	a := append(y, z...)
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	x := []int{12, 23, 34, 45, 56, 67, 78, 89, 90, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

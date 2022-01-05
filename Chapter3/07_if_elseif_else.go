package main

import "fmt"

func main() {
	x := 16
	if x < 25 {
		fmt.Printf("%v is less than 25", x)
	} else if x > 25 {
		fmt.Printf("%v is greater than 25", x)
	} else {
		fmt.Printf("%v is a number", x)
	}
}

package main

import "fmt"

func main() {
	i := 1997
	for {
		fmt.Println(i)
		i++
		if i > 2022 {
			break
		}
	}
}

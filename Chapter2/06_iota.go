package main

import "fmt"

const (
	a = 2022 + iota
	b = 2023 + iota
	c = 2024 + iota
	d = 2025 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

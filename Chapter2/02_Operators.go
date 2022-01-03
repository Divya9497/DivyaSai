package main

import "fmt"

func main() {
	a := (25 == 26)
	b := (25 <= 26)
	c := (25 >= 26)
	d := (25 != 26)
	e := (25 < 26)
	f := (25 > 26)
	fmt.Println(a, b, c, d, e, f)
}

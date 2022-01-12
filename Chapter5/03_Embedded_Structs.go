package main

import "fmt"

func main() {
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourwheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		fourwheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.doors)
	fmt.Println(t1.color)
	fmt.Println(s1.doors)
	fmt.Println(s1.color)
}

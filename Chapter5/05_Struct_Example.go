package main

import "fmt"

type PVehicle struct {
	name    string
	mileage int
}
type DVehicle struct {
	name    string
	mileage int
}

func (p PVehicle) MileageKM() int {
	return p.mileage * 2
}
func (p PVehicle) ConvertToMeter() int {
	return p.mileage * 1000
}
func main() {
	p1 := PVehicle{
		name:    "Verna",
		mileage: 40,
	}
	d1 := DVehicle{
		name:    "Nexon",
		mileage: 60,
	}
	fmt.Printf("The %v vehicle mileage is %d\n", p1.name, p1.mileage)
	fmt.Printf("The %v vehicle mileage is %d\n", d1.name, d1.mileage)
	x := p1.MileageKM()
	y := p1.ConvertToMeter()
	fmt.Println(x)
	fmt.Println(y)
}

package main

import "fmt"

type car struct {
	name     string
	colour   string
	maxspeed int
	year     int
	fuel     int
}

func (r *car) Drive() {
	if r.fuel > 3 {
		fmt.Print("car is moving ")
	} else {
		fmt.Print("please fill fuel ")
	}
}

func (r *car) Colour() string {

	return r.colour
}

func (r *car) Switchon() {
	fmt.Printf("%v %v switched on !!! ", r.name, r.colour)
	r.Drive()

}
func main() {
	carr := car{
		name:     "gentra",
		colour:   "Qora",
		year:     2023,
		maxspeed: 220,
		fuel:     5,
	}

	carr.Switchon()

}

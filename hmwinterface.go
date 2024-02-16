package main

import "fmt"

/*
type Worker interface {
	work() string
}

func Printwork(k Worker) {

	fmt.Printf("your work is %v\n", k.work())

}

type workk struct {
	worrk string
}

func (l workk) work() string {

	return l.worrk

}

func main() {

	t := workk{
		worrk: "logist",
	}

	Printwork(t)

}
*/

type shape interface {
	area() float64
}

func Prinareaa(e shape) {

	fmt.Println("area of area : ", e.area())

}

type uburchak struct {
	height float64
	base   float64
}

func (u uburchak) area() float64 {

	return float64(1) / float64(2) * u.height * u.base
}

type torburch struct {
	height float64
	basee  float64
}

func (p torburch) area() float64 {

	return p.height * p.basee
}

func main() {

	uch := uburchak{
		height: 7445.4,
		base:   392.4,
	}

	area := uch.area()
	fmt.Printf("area this %5f\n", area)

	to := torburch{
		height: 45.3,
		basee:  54.3,
	}

	fmt.Println("torburchak : ", to.area())

	uuch := uburchak{
		height: 55,
		base:   54.4,
	}

	Prinareaa(&uuch)
	Prinareaa(to)

}

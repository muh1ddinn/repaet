package main

import "fmt"

/*
type Animal interface {
	MakeSound() string
	Move() string
}

func PrintSound(a Animal) {

	fmt.Printf("sound: %s,Movement: %s\n", a.MakeSound(), a.Move())

}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "woof"
}

func (d Dog) Move() string {
	return "Runing"
}

type Bird struct{}

func (b Bird) MakeSound() string {

	return "tweet"

}

func (b Bird) Move() string {

	return "flying"
}

func main() {

	dog := Dog{}
	bird := Bird{}

	PrintSound(dog)
	PrintSound(bird)

}


type Areas interface {
	area() float64
}

func Printarea(r Areas) {

	fmt.Printf("area :  %v\n", r.area())

}

type toburch struct {
	Width  float64
	Length float64
}

func (r toburch) area() float64 {

	return r.Width * r.Length
}

type circle struct {
	Radius float64
}

func (y circle) area() float64 {

	return 3.14 * y.Radius * y.Radius
}

func main() {

	toburchc := toburch{
		Width:  99,
		Length: 88,
	}

	circlee := circle{

		Radius: 99,
	}

	Printarea(toburchc)
	Printarea(circlee)
}




type books interface{

   book () string
}

func Printbook(j books) {

	fmt.Printf("book is :",j.book())
}


type bookss struct{
	author string
	title string
}

func (r book) book() string

type count int


func (d count) book()



func main(){


}*/

func PrintValue(i interface{}) {

	if v, ok := i.(string); ok {
		fmt.Println("it's a string ", v)

	} else if v, ok := i.(int); ok {
		fmt.Println("it's a int ", v)

	} else if v, ok := i.(float64); ok {
		fmt.Println("it's a int ", v)

	} else if v, ok := i.(bool); ok {
		fmt.Println("it's a bool ", v)

	} else {
		fmt.Println("UNDATA TYPE PLEASE TRY AGAIN")
	}

}

func main() {

	PrintValue(nil)

}



switch v, ok := i.(string); ok  {
case ok:
	
}
package main

import "fmt"

type person struct {
	name        string
	Scholarship int
	cours       int
}

func (p *person) courses(k *int) {
	if p.cours == 2 && p.Scholarship <= 10 {
		(*k)++
	}
}

func main() {
	var pers1, pers2, pers3, pers4, pers5 person
	var k int

	pers1.Scholarship = 25
	pers1.name = "aaaa"
	pers1.cours = 2

	pers2.name = "bbbb"
	pers2.cours = 1

	pers3.name = "cccc"
	pers3.cours = 2

	pers4.name = "dddd"
	pers4.cours = 2

	pers5.name = "ffff"
	pers5.cours = 2

	pers1.courses(&k)
	pers2.courses(&k)
	pers3.courses(&k)
	pers4.courses(&k)
	pers5.courses(&k)

	fmt.Println(k)
}

package main

import (
	"fmt"
	"time"
)

/*
func change(y, k *int) int {

	var i int

	i = *y
	*y = *k
	*k = i

	return i
}

func main() {

	var i, m int

	i = 6
	m = 7

	change(&i, &m)

	fmt.Println(i, m)

}

*/

func main() {

	var a, n int

	fmt.Print("please enter yout number which print out put = ")
	fmt.Scan(&a)
	fmt.Print("do you want how many after sec ptint to out put =")
	fmt.Scan(&n)

	if n > 3600 {

		fmt.Print("please enter another number which should be 3600 less ")
	} else {

		now := time.Now()
		for i := 0; i < a; i++ {

			fmt.Println(i)
			time.Sleep(time.Duration(time.Second * time.Duration(n)))

		}
		fmt.Print(time.Since(now))
	}
}

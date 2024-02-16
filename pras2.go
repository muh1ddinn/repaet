/*package main

import "fmt"

func sum(l [10]int) (p, s int) {

	p = 1
	for i := 1; i < 10; i++ {

		if l[i]%2 == 0 {

			p = p * l[i]
		} else if l[i]%2 != 0 {

			s = s + l[i]
		} else {
			fmt.Print("")
		}

	}
	return
}

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	p, s := sum(arr)

	fmt.Println(p)
	fmt.Println(s)


} */

package main

import (
	"fmt"
)

func swapMinMax(arr []int) int {
	minIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i

		}

	}

	f := minIndex - 0

	return f
}

func main() {
	arr := []int{9, 2, 11, 1, 5}
	k := swapMinMax(arr)

	fmt.Println(k, "ta son bor")
}

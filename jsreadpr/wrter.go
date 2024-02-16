package main

import (
	"fmt"
	"os"
	"string"
	"strings"
)

func main() {

	var n int
	fmt.Println("please enter count of people:")
	fmt.Scan(&n)

	filee, err := os.Create("people.txt")

	if err != nil {

		fmt.Println("err while creating file ")
		return
	}
	ofset := 0


	for i := 1; i <= n; i++ {

		var name, mail, job string

		fmt.Println("enter name : ")
		fmt.Scan(&name)
		fmt.Println("enter mail: ")
		fmt.Scan(&mail)
		fmt.Println("enter job : ")
		fmt.Scan(&job)


     if !(strings.HasSuffix(mail,"@gmail.com")||strings.HasSuffix(mail,"@mail.ru")){
		continue
	 }

    txt :=fmt.Sprint("%v name : %v mail: %v,job:")


	}

}

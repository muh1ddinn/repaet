package main

import "fmt"

type Branch struct {
	Id      int
	Name    string
	Address string
}

type Transaction struct {
	Id        int
	BranchId  int
	ProductId int
	Quantity  int
}

type Products struct {
	Id    int
	Name  string
	Price int
}

func main() {

	var product = []Products{
		{Id: 1, Name: "Olma", Price: 12000},
		{Id: 2, Name: "Banan", Price: 22000},
		{Id: 3, Name: "Olcha", Price: 25000},
	}
	var transaction = []Transaction{
		{Id: 1, BranchId: 1, ProductId: 1, Quantity: 12},
		{Id: 2, BranchId: 2, ProductId: 2, Quantity: 10},
		{Id: 3, BranchId: 1, ProductId: 3, Quantity: 8},
		{Id: 4, BranchId: 2, ProductId: 1, Quantity: 14},
		{Id: 5, BranchId: 1, ProductId: 2, Quantity: 13},
		{Id: 6, BranchId: 2, ProductId: 3, Quantity: 7},
	}
	var branches = []Branch{
		{Id: 1, Name: "MGorkiy", Address: "Mirzo Ulug'bek 17 uy"},
		{Id: 2, Name: "Chilonzor", Address: "Chilonzor Metro"},
	}

	var sumArr []int
	for _, k := range branches {
		sum := 0
		for _, j := range transaction {
			if j.BranchId == k.Id {
				for _, l := range product {

					if k.Id == j.BranchId && l.Id == j.ProductId {

						p := j.Quantity * l.Price

						sum = sum + p

					}
				}

			}

		}

		sumArr = append(sumArr, sum)
	}
	// Bubble sort the sumArr slice in descending order
	for i := 0; i < len(sumArr)-1; i++ {
		for j := 0; j < len(sumArr)-i-1; j++ {
			if sumArr[j] > sumArr[j+1] {
				sumArr[j], sumArr[j+1] = sumArr[j+1], sumArr[j]
			}
		}
	}

	// Print the sorted sumArr
	fmt.Println(" sums of sales for much:", sumArr[len(sumArr)-1])
}

/*
	var proidcal = 2
	var totq int

	for _, l := range transaction {
		if l.ProductId == proidcal {
			totq += l.Quantity
		}

	}

	fmt.Println("Total quantity for product with ID", proidcal, ":", totq)

*/

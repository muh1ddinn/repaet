package main

import "fmt"

type Customer struct {
	Name string
	Cash int
	Basket Basket

}

func NewCustomer(name string, cash int,basket Basket)Customer{

return Customer{

	Name: name,
	Cash: cash,
	Basket: basket,
}
}






type Basket struct {
	ProductList []Product
	Total int
}

func NewBasket() Basket{
	return Basket{}
}

type Store struct {
}

var (
	List = ProductList{
			{
			Name:          "bread",
			Count:         15,
			Price:         3000,
			OriginalPrice: 4000,
		},
		{
			Name:          "meat",
			Count:         5,
			Price:         120000,
			OriginalPrice: 90000,
		},
		{
			Name:          "milk",
			Count:         8,
			Price:         30000,
			OriginalPrice: 25000,
		},
		{
			Name:          "juice",
			Count:         20,
			Price:         13000,
			OriginalPrice: 11000,
		},
	}
)

type Product struct {
	Name          string
	Count      int
	Price         int
	OriginalPrice int
}


type ProductList []Product




func NewProduct(name string,price,count int)Product{
	return Product{
 
		Name: name,
		Price: price,
		Count: count,


	}
}

func (pr*ProductList) AddProduct(product Product){

	*pr =append(*pr,product)

}

type ProductSellRequest struct {
	Name string
	Count int
}

func GetProductinfo() ProductSellRequest{


	var ProductName string
	var count = 0


	fmt.Print("enter product name :")
	fmt.Scan(& ProductName )
	
	fmt.Print("enter product name :")
	fmt.Scan(& ProductName )
	fmt.Println()

	return ProductSellRequest{

		Name: ProductName,
		Count: count,
	}
}



main(){

const Password = "123"

const(

	Startshopcmd = iota+1
	Finishshopcmd
)


const(
   
	Bosscmd=iota+1
	Customercmdu
	Ursequitcmd
)


const (

	reportcmd=iota+1
	Listcmd
	Backcmd
	Quitcmd
)

for true{

var usercmd int  

fmt.Print("

    Enter command:
	1-Boss
	2-Customer
	3-Quit
")

fmt.Scan(&usercmd)

switch usercmd{

case Bosscmd:
	var password string 
	fmt.Println("ENTER PASSWORD : ")
	fmt.Scan(&password)
	if password != Password {

		fmt.Println("password is worng")

		return
	}

    ``

}



}



}
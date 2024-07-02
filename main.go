package main

import"fmt"

type Investment struct {

	name string
	symbol string
	quantity int
	price float64
}

func addInvestment(investments*[] Investment) {

	var name string
	var symbol string
	var quantity int
	var price float64

	fmt.Println("Please enter the investment name:")
	fmt.Scan(&name)

	fmt.Println("Please enter the investment symbol:")
	fmt.Scan(&symbol)

	fmt.Println("Please enter the investment quantity:")
	fmt.Scan(&quantity)

	fmt.Println("Please enter the investment price:")
	fmt.Scan(&price)

	var inv Investment

	inv.name = name
	inv.symbol = symbol
	inv.quantity = quantity
	inv.price = price

	*investments = append(*investments,inv)
	

}

func updateInvestmentQuantity(investments*[] Investment){

	var name string
	var quantity int

	fmt.Println("Which investment do you want to update?")

	fmt.Scan(&name)

	fmt.Println("What is the quantity:")

	fmt.Scan(&quantity)

	for i:=0; i < len(*investments); i++ {
		if(*investments)[i].name == name {
			(*investments)[i].quantity = quantity
			break
		}
	}
}

func calculateTotalValueOfPortofolio(investments[] Investment)float64{

	var total float64 
	

	for i:=0; i < len(investments); i++ {
		total += investments[i].price * float64(investments[i].quantity)
		
	}
	return total
		
}

func listOfAllInvestments(investments[] Investment) {

	for _, investment := range investments {
	fmt.Println("Name:", investment.name, "| Symbol:", investment.symbol, "| Quantity:", investment.quantity, "| Price:", investment.price)
	}



}


func main() {

var investments = [] Investment{}

running := true

var menuOption int

for running == true {

	fmt.Println("1 - Add investment")
	fmt.Println("2 - Update investment quantity")
	fmt.Println("3 - Calculate Total Value of Portofolio")
	fmt.Println("4 - List Investment")
	fmt.Println("5 - Exist")

	fmt.Println("Please enter the menu option:")
	fmt.Scan(&menuOption)

	if menuOption == 1 {
		addInvestment(&investments)
	} else if menuOption == 2 {
		updateInvestmentQuantity(&investments)
	} else if menuOption == 3 {
		var total float64 = calculateTotalValueOfPortofolio(investments)
		fmt.Printf("The total value of the portofolio is %.2f", total)
	} else if menuOption == 4 {
		listOfAllInvestments(investments)
	} else if menuOption == 5 {
		running = false
	} else {
		fmt.Println("Invalid option!!!")
	}
	} 

}






package main 

import"fmt"

/*

SuperMarket project

This SuperMarket project program start by askig the user to insert product data (id, name,price,stock)
this entering of the product data stops only when the user enters -1 to the id.

Then the user will be presented a menu with 4 option: Remove Product, List All Product and Exit.



*/

func getProductsData(ids*[]int, names*[]string, prices*[]float64, stocks*[]int ){

	running := true

	var id int
	var name string
	var price float64
	var stock int

	for running == true {


		fmt.Println("Please enter the product's id:")
		fmt.Scan(&id)

		if id < 1 && id != -1 {
			continue
		} else if id == -1 {
			running = false
			continue
		}

		fmt.Println("Please enter the product's name:")
		fmt.Scan(&name)

		fmt.Println("Please enter the product's price:")
		fmt.Scan(&price)

		fmt.Println("Please enter the product's stock")
		fmt.Scan(&stock)

		*ids = append(*ids, id)

		*names = append(*names, name)

		*prices = append(*prices, price)

		*stocks = append(*stocks, stock)
	}



}

func addProduct(ids*[] int, names*[]string, prices*[] float64, stocks*[] int){

	var id int = 0
	var name string
	var price float64 = 0
	var stock int = 0 

	for id <= 0 {
		fmt.Println("Please enter the product's id:")
		fmt.Scan(&id)	
		
		for i:=0; i < len(*ids); i ++ {
			if (*ids)[i] == id {
				fmt.Println("Error: Product ID saved can't be reused!")
				return
			}
		}
	}

	fmt.Println("Please enter the product's name:")
	fmt.Scan(&name)

	for price <= 0{
		fmt.Println("Please enter the product's price:")
		fmt.Scan(&price)
	}

	for stock <=0 {
		fmt.Println("Please enter the product's stock")
		fmt.Scan(&stock)

		*ids = append(*ids, id)

		*names = append(*names, name)

		*prices = append(*prices, price)

		*stocks = append(*stocks, stock)
	}

}

func removeProduct(ids*[] int, names*[]string, prices*[] float64, stocks*[] int){

	var id int
	fmt.Println("Please enter the id of the product you want to remove: ")
	fmt.Scan(&id)
	for i:=0; i<len(*ids); i++ {
		if id == (*ids)[i] {
		   copy((*ids)[:i], (*ids)[i+1:])
		   *ids = (*ids)[:len(*ids)-1]

		   copy((*names)[:i], (*names)[i+1:])
		   *names = (*names)[:len(*names)-1]

		   copy((*prices)[:i], (*prices)[i+1:])
		   *prices = (*prices)[:len(*prices)-1]

		   copy((*stocks)[:i], (*stocks)[i+1:])
		   *stocks = (*stocks)[:len(*stocks)-1]

		   break

		}
	}

}

func calculateAveragePrice(prices [] float64) float64{

	var sum float64 = 0

	for i:=0; i < len(prices); i++ {
		sum += prices[i]
	}

	return sum/float64(len(prices))
}

func calculatePricesStandardDeviation( prices[] float64) {

var sum float64 = 0 

for i := 0; i < len(prices); i ++ {
	sum += prices[i]

}

//avg := sum/float64(len(prices))



}


func listAllProducts(ids[] int, names[] string,  prices[] float64, stocks[] int ) {

	for i:=0; i < len(ids); i++ {
		fmt.Println("| Name:", names[i], "| Prices:",prices[i], "|Stock:", stocks[i])
	}


}

func main(){

	var ids = []int{}
	var names = []string{}
	var prices = []float64{}
	var stocks = []int{} 

	getProductsData(&ids, &names, &prices, &stocks)

	var userInput int = 0

	for true {

	fmt.Println("----------------MARIANA SUPERMARKET-------------------")

	fmt.Println("1 - Add Product")
	fmt.Println("2 - Remove Product")
	fmt.Println("3 - List All Products")
	fmt.Println("4 - Exit")

	fmt.Println("Please choose an option: ")
	fmt.Scan(&userInput)

	if userInput == 1 {
		addProduct(&ids, &names, &prices, &stocks)
	} else if userInput == 2 {
		removeProduct(&ids, &names, &prices, &stocks)
	} else if userInput == 3 {
		listAllProducts(ids, names, prices, stocks)
	} else if userInput == 4 {
		break
	}
}

	//avg := calculateAveragePrice(prices)

	//calculatePricesStandardDeviation()

	//print(avg)
	
	
	/*for i:=0; i < len(ids); i++ {
		fmt.Println("Id:",ids[i], "| Name:", names[i], "| Prices:",prices[i], "|Stock:", stocks[i])
	} */
	
}
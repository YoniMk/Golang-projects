package main

import"fmt"


/*

Create a project that asks for student's name, age, yearGrouo, gradesAverage

*/

/*

Sruct is the possibilty for us programmers to create our own data type

*/

type student struct {

	name string
	age int
	yearGroup string
	gradeAverage float64

}

type product struct {
	id int
	name string
	price float64
	stock int 
}

type footballPlayer struct {
	name string
	team string
	shirtNumber int
	marketValue float64
	goalScored int

}

func main() {

	var student1 student

	student1.name = "Yoni"
	student1.age = 28
	student1.yearGroup = "1st grade"
	student1.gradeAverage = 13.7

	var player footballPlayer

	player.name = "Ronaldinho"
	player.team = "Barcelona"
	player.shirtNumber = 10
	player.marketValue = 200000
	player.goalScored = 20

	var products = []product{}

	productsCounter := 0

	for i:=0; i< 5; i++ {

		productsCounter += 1
		products[i].id = productsCounter
		fmt.Println("Please enter the product name: ")
		fmt.Scan(&products[i].name)
		fmt.Println("Please enter the product price: ")
		fmt.Scan(&products[i].price)
		fmt.Println("Please enter the product stock: ")
		fmt.Scan(&products[i].stock)
	}



}
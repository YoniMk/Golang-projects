package main



/*

Build a project that asks the user for positive integer numbers and accumulates them
on a variable.
The program will end when the user enters a negative number

*/
import "fmt"

func main () {

var number int = 1
var sum int

for number > 0 {
	fmt.Println("Please enter a positive number: ")
	fmt.Scan(&number)
	if number > 0 {
		sum += number // sum = sum + number	
	} else {
		fmt.Println("The number you entered is negative, the program will finish!")	
    }
}

   fmt.Println("The final sum is:", sum)

}
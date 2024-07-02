package main

import "fmt"
import "bufio"
import "os"

func main() {

//var name string = "Gabriella"
reader := bufio.NewReader(os.Stdin)

 names := [5]string{"Gabriella", "Lorenzo", "Yoni", "Sylvie", "Alfred"}

 var numbers = [10]int{}

 for j:=0; j < len(numbers); j++ {

	fmt.Println("Please enter a number")

	fmt.Scan(&numbers[j])
 }

 for i := 0; i< len(names); i++ {
	fmt.Println(names[i])
 }

 var favoriteMusics = [5]string{}

 for j:=0; j < len(favoriteMusics); j++ {

	fmt.Println("please enter the name of a song:")



	favoriteMusics[j] ,_ = reader.ReadString('\n')
 }

 for j := 0; j< len(favoriteMusics); j++ {
	fmt.Println(favoriteMusics[j])
 }

favoriteMusics[0] = "undefined"

fmt.Println(favoriteMusics[0])


var slice1 = numbers[0:4] // position 4 is not included

for i:=0; i < 4; i++ {
	slice1[i] = 20
}


for j:=0; j < len(numbers); j ++ {

	fmt.Printf("Number in position %d: %d\n", j, numbers[j])
	
}

var s = append(slice1, 20) //adding another 20 to the slice

fmt.Println(s)

fmt.Println(numbers)

for index, a := range numbers {
	fmt.Println(index, ":", a)
}

}
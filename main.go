package main 

import "fmt"

/*
	ARRAYS VS SLIVES

*/

func main() {

	/*var grades = [10]int{ }

	for i:=0 ; i < len(grades); i ++ {
		fmt.Println("please enter a student grade:")
		fmt.Scan(&grades[i])
	}


	// How can I add one more grade to this array?

	var grades2 = [11]int{}

		for i:=0; i < len(grades); i++ {
			grades2[i] = grades[i]
		}

		grades2[10] = 20

		//How can I remove a grade from the array?
        
		grades[9] = -1

		for i:=0; i < len(grades); i++ {
			if grades[i] < 0 {
				fmt.Println("Please enter a grade:")
				fmt.Scan(&grades[i])
			}
		}

		for i:=0; i < len(grades); i++ {
				fmt.Printf("Grade %d : %d\n", i, grades[i])
		}

		
		//g = append(g,1) // adding the number 1 to my slice "g" will now have size 1
		var num int

		for i:=0; i < 10; i ++ {
			fmt.Println("Please enter a student grade:")
			fmt.Scan(&num)
			g = append(g, num)
		}

		*/
		run := true
		num := 0
		var g = []int{} //slices have no fixed size


		for run == true {
			fmt.Println("Please enter a student grade:")
			fmt.Scan(&num)
			if num < 0 {
				run = false
			} else {
				g = append(g, num)
			}
		}



//We want to remove the firsr grade

g[0] = g[len(g) - 1] // copying the grade of the last position to the first position.

g = g[0:len(g)-1] // goes from 0 to 4, 4 not included, stops in 3.

for i:=0; i< len(g); i++ {
	fmt.Println(g[i])
}
	
}
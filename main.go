package main

import(
	 "fmt"
	 "os"
	 "encoding/csv"
	 "strconv"
)

type Student struct {
	loginName string
	firstName string
	familyName string
	level int
	class string
	bestResult int

}

func main() {

	file, err := os.Open("students.csv")

	students := []Student{}

	if err != nil { //If there's an error
		fmt.Println("Error opening the file")
		return
	}

	reader := csv.NewReader(file) //this is not reading the file, it's just creating the reader

	reader.Read()

	for { //infinite loop, only ends if you do a break

		line, err := reader.Read()

		if err != nil {
			break
		}

		var newStudent Student

		newStudent.loginName = line[0]
		newStudent.firstName = line[1]
		newStudent.familyName = line[2]
		newStudent.level, err = strconv.Atoi(line[3])
		if err != nil {
			fmt.Println("Their was a problem in conversion", err)
			newStudent.level = 0
		}
		newStudent.class = line[4]
		newStudent.bestResult, err = strconv.Atoi(line[5])

		if err != nil {
			fmt.Println("Their was a problem in conversion", err)
			newStudent.bestResult = 0
		}
		
		students = append(students, newStudent)


	}

	for _,student := range students {
		fmt.Printf("The student %s at level %d was imported from the csv file to the program\n", student.loginName,student.level)
	}



}
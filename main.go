package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//The main function is the entry point of the program. It's where execution begins
func main() {
	
	//Relational Operators
	/*
	== Equal
	!= Not Equal
	< less than
	> greater than
	<= less than or equal
	>= greater or equal
	*/

	age := 19
	
	fmt.Println("Is age equal to 23:",age == 23)
	fmt.Println("Is age not equal to 23:",age != 23)
	fmt.Println("Is age less than  23:",age < 23)
	fmt.Println("Is age greater than  23:",age > 23)

	stage := "TD"

	evaluation := stage == "TD" && 10 < 18
	evaluation2 := stage == "TD" || 10 < 18
	myBoolean := false
	invertedBoolean := !myBoolean

	fmt.Println(evaluation)
	fmt.Println(evaluation2)
	fmt.Println(invertedBoolean)
}

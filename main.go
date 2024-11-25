package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//The main function is the entry point of the program. It's where execution begins
func main() {
	//? ARRAYS

	// var integers[5] int // [0,0,0,0,0] initial values

	//We can modifiy its values
	//Has to be same type
	// integers[0] = 100
	// integers[1] = 200
	// integers[2] = 300
	// integers[3] = 400
	// integers[4] = 500

	// fmt.Println(integers)

	//OTHER WAY TO CREATE ARRAYS

	// numbers := [3] int {100,200,300}
	// numbers2 := [...] int {100,200,300} // [...] Go will automatically detect the lenght

	// fmt.Println(numbers)
	// fmt.Println(numbers2)

}

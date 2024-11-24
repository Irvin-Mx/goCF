package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//Constants outside of 
const speed int = 100
const message = "Warning!"


//The main function is the entry point of the program. It's where execution begins
func main() {
	
	fmt.Println(speed,message)
}

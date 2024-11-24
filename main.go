package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//The main function is the entry point of the program. It's where execution begins
func main() {
	//? In Go there are many ways to manage standard input, here is one example
	//! fmt.Scanf is limited to only one string with no spaces, consider
	var sn int 
	var opName string
	var pn float32
	var leds int

	fmt.Print("Enter the Serial Number of the card: SN")
	fmt.Scanf("%d", &sn) // int value, variable

	fmt.Print("Enter the leds that light up:")
	fmt.Scanf("%s", &leds)  // string value, variable

	fmt.Print("Enter your opName")	
	fmt.Scanf("%s", &opName)  // string value, variable

	fmt.Print("Enter the pn:")
	fmt.Scanf("%f", &pn)  // string value, variable


	fmt.Printf("Part Number: %.2f Leds: %d opName %s SN: %d", pn,leds,opName,sn)
}

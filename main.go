package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//The main function is the entry point of the program. It's where execution begins
func main() {
	var isHigh string // empty string "" initial value
	var pressure int // initial value 0
	var frequency = 32.32
	otherVariable := true

	var temp int = 99

	var maxTemp int
	maxTemp = 90
	if temp > maxTemp {
		fmt.Println("The temperature is too high")
	}
	for i := 0; i < 10; i++ {
		temp ++
	}
	fmt.Println(isHigh)
	fmt.Println(pressure)
	fmt.Println(temp)
	fmt.Println(otherVariable)
	fmt.Println(frequency)
	fmt.Println("Is high:",isHigh,"Pressure:",pressure,"Temp:",temp,"Other Variable:",otherVariable,"Frequency:",frequency)
}

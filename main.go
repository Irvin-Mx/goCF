package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//The main function is the entry point of the program. It's where execution begins
func main() {
	/* var productName string
	 var productPrice float32
	 var productSku string
	 var productNumber int
	*/

	// var productName,productPrice,productSku,productNumber = "a","b","c","d"
	// productName,productPrice,productSku,productNumber := "a","b","c","d"
	// item1,item2,item3 := "x",1,true

	// Variables of the same type
	 var productName,productPrice,productSku,productNumber string
	productName = "GoPro"
	productPrice = "5,700"
	productSku = "323242aa3523124s44f2"
	productNumber = "324353"
	fmt.Println(productName,productNumber,productPrice,productSku)
}

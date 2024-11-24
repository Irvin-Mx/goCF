package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O
import "reflect"
//The main function is the entry point of the program. It's where execution begins
func main() {
	// var url string = "www.someurl.com"
	// var url = "string"
	url := "http://www.jsonplaceholder.com/todos"

	fmt.Println(url)
	fmt.Println(len(url))
	fmt.Println(url[0]) //* ASCII

	//We use reflect package
	fmt.Println(reflect.TypeOf(url[0]))

	someCharacter := url[14]

	fmt.Printf("%c\n", someCharacter)
}

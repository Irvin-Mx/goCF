package main // name of our package

import "fmt" //Imports the formatting package from Go's standard library, which provides functions for formatted I/O

//Secuencial constants example
// const a int = 0
// const b int = 1
// const c int = 3
// const d int = 4
// const e int = 5
// const f int = 6
// const g int = 7
//*Refactored below
// const (
// 	a int = 0
// 	b int = 1
// 	c int = 3
// 	d int = 4
// 	e int = 5
// 	f int = 6
// 	g int = 7
// )

//*Iota example
 const (
 	a int = iota * 10 // Will start at 0, if needed can make other operations like +,-,*,/
 	b
 	c
 	d
 	e
 	f
 	g
 )


//The main function is the entry point of the program. It's where execution begins
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

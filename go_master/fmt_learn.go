package main

import "fmt"

var s string = `I am a variable`
var i int = 24

func main() {
	fmt.Println("This is start of main block")
	foogen()
	foobol()
	fooint()
}

func foogen() {
	fmt.Println("Notes -- Printing the variable value")
	fmt.Println(s)
	fmt.Println("Notes -- Printing the variable's type")
	fmt.Printf("%T\n", s)
	fmt.Println("Notes -- Printing the value in default format")
	fmt.Printf("%v\n", s)
	fmt.Println("Notes -- Printing the Go-syntax representation of the value")
	fmt.Printf("%#v\n", s)
	fmt.Println("Notes -- a literal percent sign. consumes no value")
	fmt.Printf("%%\n")
}

func foobol() {
	fmt.Printf("%t\n", true)
}

func fooint() {
	fmt.Println(i)
}

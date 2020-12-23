package main

import "fmt"

type app []string

type Munchin interface{}

func WhichType(x Munchin) {
	switch x.(type) {
	case (string):
		fmt.Println("This is string type")
	case (int):
		fmt.Println("This is int type")
	case ([]string):
		fmt.Println("This is slice of strings")
	case ([]int):
		fmt.Println("This is slice of ints")
	default:
		fmt.Println(`This is not one of "string","int","string slice" or "int slice". Might be something else`)
	}
}

func main() {
	a := app{"Bhuvanesh", "manne"}
	WhichType(a)
}
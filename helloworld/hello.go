package main

import "fmt"

func main() {
	fmt.Println("Hello; This is Bhuvan")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo")
}

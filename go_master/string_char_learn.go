package main

import "fmt"

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {

	for i := 0; i < len(sample); i++ {
		slice := sample[i:len(sample)]
		fmt.Println("Println:")
		fmt.Println(slice)
		fmt.Printf("%T\n", slice[i])
		fmt.Printf("%x\n", slice)
		fmt.Println("Printf with % q:")
		fmt.Printf("%q\n", slice)

		fmt.Println("Printf with % +q:")
		fmt.Printf("%+q\n", slice)
		fmt.Println("==================================")
	}
}
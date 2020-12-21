package main

import (
	"fmt"
)

func main() {

	sa := "ABCDEFGH"
	a := []byte(sa)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("%s\n", a)
	kr := other(a)
	fmt.Printf("%s", kr)
}

func other(a []byte) []byte {
	i, j := 0, len(a)-1
	for {
		if i < j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
	return a
}

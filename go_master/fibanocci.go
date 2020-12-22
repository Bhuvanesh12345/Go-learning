package main

import "fmt"

func main() {
	res := fibb(100)
	fmt.Println(res)
}

func fibb(n int) []int {
	bnum := []int{0, 1}
	for i := 0; i <= n; i++ {
		t := len(bnum)
		sval := bnum[t-1:][0] + bnum[t-2 : t-1][0]
		bnum = append(bnum, sval)
	}
	return bnum
}

package main

import "fmt"

type person struct {
	first string
	last  string
	icl   []string
}

var l []person

func main() {

	t1 := addp("Prem", "Kumar", []string{"choco", "Blackber", "cream"}, l)
	t2 := addp("Prudhvi", "Manne", []string{"chocolate", "marshmallow", "lichi"}, t1)
	t3 := addp("Bhuvanesh", "Manne", []string{"Butterscotch", "Vanilla"}, t2)
	for _, v := range t3 {
		fmt.Println(v.first, v.last)
		for j, w := range v.icl {
			fmt.Println("\t\t", j, w)
		}
	}
}

func addp(first_name string, last_name string, favs []string, l []person) []person {
	k := person{
		first: first_name,
		last:  last_name,
		icl:   favs,
	}
	lcl := append(l, k)
	return lcl
}

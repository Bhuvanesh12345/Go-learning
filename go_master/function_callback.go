package main

import (
	"fmt"
)

func main() {
	//xi := genseq(10, 1000)
	//k := sum(xi)
	//fmt.Println(k)
	av, ev := even(genseq, 1, 100)
	fmt.Printf("The generated even number sequence is as below:\n%v\n", av)
	fmt.Printf("The total of this sequence is :%v", ev)
}

func sum(p []int) int {
	total := 0
	for _, v := range p {
		total += v
	}
	return total

}

func genseq(p, q int) []int {
	var tp []int
	for i := p; i <= q; i++ {
		tp = append(tp, i)
	}
	return tp
}

//generate even numbers and sum
func even(f func(p, q int) []int, p int, q int) ([]int, int) {
	tp := f(p, q)
	var kp []int
	for _, v := range tp {
		if v%2 == 0 {
			kp = append(kp, v)
		}
	}
	sp := sum(kp)
	return kp, sp
}

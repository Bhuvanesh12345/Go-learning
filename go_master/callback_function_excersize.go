package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius

}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

// Function returns a function
func info(s shape) func() float64 {
	k := s.area()
	return func() float64 {
		return k
	}
}

//Passing fucnction into a function as argument (callback)
//Checking if area is always positive
func checker(a func() float64) bool {
	if a() >= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	c := circle{
		radius: 4,
	}
	s := square{
		length: 4,
	}
	fmt.Println(checker(c.area))
	fmt.Println(checker(s.area))
	func() {
		fmt.Printf("This is the length of the square: %v and area of the square: %v\n", s.length, info(s)())
		fmt.Printf("This is the radius of the circle: %v and area of the circle: %v\n", c.radius, info(c)())
	}()

}

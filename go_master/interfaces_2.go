package main

import (
	"fmt"
)

type animal struct {
	name     string
	location string
}

type domer interface {
	domorwild() animal
}

func (a animal) domorwild() animal {
	switch a.name {
	case "lion":
		fmt.Println(" I am lion, I live in jungle and I am wild")
	case "calf":
		fmt.Println(" I am calf, I live near humans and I am domestic")
	default:
		fmt.Println("Printing something")
	}
	return a
}

func main() {
	lion := animal{
		name:     "lion",
		location: "jungle",
	}
	calf := animal{
		name:     "calf",
		location: "home",
	}
	whatIeat(lion)
	whatIeat(calf)

}

func whatIeat(d domer) {
	if d.domorwild().location == "jungle" {
		fmt.Println("My functionality to eat:  I eat fellow animals in jungle")
	} else {
		fmt.Println("My functionality to eat:  I eat leaves")
	}
}

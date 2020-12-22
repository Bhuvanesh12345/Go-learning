package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"Bhuvanesh":  "Leader",
		"Friends":    "Members",
		"Relatives":  "Standby",
		"Neighbours": "Nothing",
	}

	m["something else"] = 33
	delete(m, "Bhuvanesh")
	for val, _ := range m {
		if i, v := m[val]; v {
			fmt.Printf("%v is my relationship with my %v\n", i, val)
		}
	}

}

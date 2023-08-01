package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range num {
		if val%2 == 0 {
			fmt.Printf("Even no: %v\n", val)
		} else {
			fmt.Printf("Odd no: %v\n", val)
		}
	}
}

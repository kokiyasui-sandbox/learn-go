package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 100
	arr := make([]int, n)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	for _, v := range arr {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Printf("%d, Six!\n", v)
		case v%2 == 0:
			fmt.Printf("%d, Two!\n", v)
		case v%3 == 0:
			fmt.Printf("%d, Three!\n", v)
		default:
			fmt.Printf("%d, Never Mind!\n", v)
		}
	}
}

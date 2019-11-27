package main

import "fmt"

func main() {
	// for init; condition; post{}
	for i := 0; i <= 3-1; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j <= 4-1; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}
}

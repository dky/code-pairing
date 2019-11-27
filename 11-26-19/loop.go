package main

import "fmt"

func main() {
	// for init; condition; post{}
	for i := 0; i <= 3-1; i++ {
		for j := 0; j <= 2-1; j++ {
			fmt.Println(j)
		}
	}
}

package main

import "fmt"

func main() {
	a := []int{1, 3, 2}

	for i := 0; i < len(a)-1; i++ {
		//fmt.Println(i)
		for j := 0; j < len(a); j++ {
			fmt.Println(i, j)
		}
	}
}

package main

import "fmt"

func main() {
	list := []int{4, 3, 2, 1}

	for i, v := range list {
		fmt.Println(i, v)

	}
}

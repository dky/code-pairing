package main

import "fmt"

func main() {
	toSort := [5]int{50, 40, 20, 30, 10}

	sorted := bubbleSort(toSort)

	incorrectSort(toSort)

	fmt.Println("Initial array:", toSort)
	fmt.Println("Sorted array:", sorted)

}

func bubbleSort(input [5]int) [5]int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				//fmt.Println(i, j)
				//fmt.Println(input[i], input[j])
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
				//fmt.Println("Swap")
				//fmt.Println(input[i], input[j])
			}
		}
	}
	//fmt.Println("input array")
	//fmt.Println(input)

	return input
}

func incorrectSort(input [5]int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i] > input[j] {
				//fmt.Println(i, j)
				//fmt.Println(input[i], input[j])
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
				//fmt.Println("Swap")
				//fmt.Println(input[i], input[j])
			}
		}
	}
	//fmt.Println("input array")
	fmt.Println("Incorrect array:", input)

	//return input
}

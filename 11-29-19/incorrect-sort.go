package main

import "fmt"

// bubbleSort takes in a slice of ints
func bubbleSort(input []int) {
	fmt.Printf("Number of elements in slice: %d\n", len(input))
	fmt.Printf("Unsorted slice: %d\n\n", input)

	//Start from index position 1
	//dky why do we start from 1 vs 0?
	for i := 1; i < len(input); i++ {
		//Start from index position 0
		for j := 0; j < len(input); j++ {
			//fmt.Printf("index i: %d, element i: %d index j: %d, element j: %d\n", i, input[i], j, input[j])
			fmt.Printf("index i: %d, index j: %d\n", i, j)
			fmt.Printf("element i: %d element j: %d\n", input[i], input[j])
			order(input[i], input[j], input)
		}
	}
}

// order takes in the input[i] and input[j] and swaps their place values if i > j
func order(i, j int, inputSlice []int) []int {
	fmt.Printf("comparing element i: %d and element j: %d\n", i, j)
	if i > j {
		fmt.Println("i is greater than j, swapping values")

		tmp := i //tmp stores the value of i
		i = j    //assign j to i
		j = tmp  //assign tmp to j

		fmt.Printf("i: %d, j: %d\n", i, j)
		fmt.Println(inputSlice)
		fmt.Printf("\n")
	} else {
		fmt.Println("i is not greater than j, going back to top of loop j")
		fmt.Println(inputSlice)
		fmt.Printf("\n")
	}

	return inputSlice
}

func main() {
	unsortedList := []int{30, 10, 20}
	//unsortedList := []int{30, 10, 20, 40}
	//unsortedList := []int{50, 40, 20, 30, 10}
	bubbleSort(unsortedList)
}

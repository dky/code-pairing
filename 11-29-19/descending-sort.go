package main

import "fmt"

// printSlice returns the current state of the working slice.
func printSlice(currentSlice []int) {
	fmt.Println("current state of slice:", currentSlice)
}

// bubbleSort takes in a slice of ints
func bubbleSort(input []int) {

	//Start from index position 1
	//dky why do we start from 1 vs 0?
	for i := 1; i < len(input); i++ {
		//Start from index position 0
		for j := 0; j < len(input); j++ {
			fmt.Printf("Number of elements in slice: %d\n", len(input))
			fmt.Printf("Unsorted slice: %d\n", input)
			fmt.Printf("index i: %d, index j: %d\n", i, j)
			fmt.Printf("element i: %d element j: %d\n", input[i], input[j])
			order(i, j, input)
		}
	}
}

// order takes in the input[i] and input[j] and swaps their place values if i > j
func order(i, j int, input []int) {

	fmt.Printf("comparing element i: %d and element j: %d\n", input[i], input[j])

	if input[i] > input[j] {
		fmt.Println("element i is greater than element j, swapping values")

		tmp := input[i]   //tmp stores the value of i
		swapi := input[j] //assign j to i
		swapj := tmp      //assign tmp to j
		fmt.Printf("new values for i: %d and j: %d\n", swapi, swapj)

		fmt.Printf("replacing elements at input[i]: %d, and input[j]: %d\n", input[i], input[j])
		input[i] = swapi
		input[j] = swapj
		printSlice(input)

		fmt.Printf("\n")
	} else {
		fmt.Println("i is not greater than j, going back to top of loop j")
		printSlice(input)
		fmt.Printf("\n")
	}
}

func main() {
	//unsortedList := []int{35, 12, 40} //3 element list
	//unsortedList := []int{30, 10, 20, 40} //4 element list
	unsortedList := []int{50, 40, 20, 30, 10} //5 element list
	bubbleSort(unsortedList)
}

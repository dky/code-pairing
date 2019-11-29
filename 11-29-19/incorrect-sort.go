package main

import "fmt"

func main() {
	//unsorted := []int{30, 20, 10}
	unsorted := []int{50, 40, 20, 30, 10}
	bubbleSort(unsorted)
}

// The issue with incorrectSort is it returns the sort in decending order. Why?
func bubbleSort(input []int) {
	fmt.Println("Number of elements in slice:", len(input))
	//Start from index position 1
	for i := 1; i < len(input); i++ {
		//Start from index position 0
		for j := 0; j < len(input); j++ {
			fmt.Printf("index i: %d, element i: %d index j: %d, element j: %d\n", i, input[i], j, input[j])
			//order(input[i], input[j])
		}
	}
	fmt.Println("Sorted output:", input)
}

// order takes in the input[i] and input[j] and swaps their place values if they are in the wrong place.
func order(i, j int) {
	if i > j {
		// Create a tmp var to store the value of i
		tmp := i
		i = j
		j = tmp
		fmt.Printf("Comparing i: %d, j: %d\n", i, j)
	}
}

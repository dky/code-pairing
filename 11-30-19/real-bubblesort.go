package main

import "fmt"

func main() {
	// some tests
	unsortedSlice1 := []int{20, 3, 5, 1, 40, 3}
	unsortedSlice2 := []int{}
	unsortedSlice3 := []int{1, 2, 3}

	fmt.Println(bubbleSort(unsortedSlice1))
	fmt.Println(bubbleSort(unsortedSlice2))
	fmt.Println(bubbleSort(unsortedSlice3))
}

// this is the real bubble sort algorithme , sorry my bad... it's super similar to the one we did though
func bubbleSort(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		isSwapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
				isSwapped = true
			}
		}

		if !isSwapped {
			return arr
		}
	}

	return arr
}

func swap(left *int, right *int) {
	tmp := *left
	*left = *right
	*right = tmp
}

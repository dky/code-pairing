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

func bubbleSort(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		isSwapped := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
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

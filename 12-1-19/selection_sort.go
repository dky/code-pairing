package main

import "fmt"

// Adapting this from https://www.youtube.com/watch?v=GUDLRan2DWM

func main() {
	arrayA := []int{3, 1, 5, 2, 4}

	min := arrayA[0]
	sliceLength := len(arrayA)
	//is 8 < 8? n
	//is 7 < 8? y min = 7

	for i := 1; i < sliceLength; i++ {
		fmt.Printf("Loop count: %d\n", i)
		if arrayA[i] < min {
			fmt.Printf("Is %d < %d?\n", arrayA[i], min)
			fmt.Printf("yes, min: %d\n", arrayA[i])
			min = arrayA[i]

			//arrayA[0] = min
			for j := i + 1; j < sliceLength; j++ {

				fmt.Println(j)

			}

		} else {
			fmt.Printf("%d > %d\n", arrayA[i], min)
		}
	}

	//fmt.Println(min)
	fmt.Println(arrayA)

}

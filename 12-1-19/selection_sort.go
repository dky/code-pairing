package main

import "fmt"

// Adapting this from https://www.youtube.com/watch?v=GUDLRan2DWM

//func main() {
//arrayA := []int{3, 1, 5, 2, 4}

//min := arrayA[0]
//sliceLength := len(arrayA)
////is 8 < 8? n
////is 7 < 8? y min = 7

//for i := 1; i < sliceLength; i++ {
//fmt.Printf("Loop count: %d\n", i)
//if arrayA[i] < min {
//fmt.Printf("Is %d < %d?\n", arrayA[i], min)
//fmt.Printf("yes, min: %d\n", arrayA[i])
//min = arrayA[i]

////arrayA[0] = min
////for j := 0; j < sliceLength; j++ {
////for j := i; j < sliceLength; j++ {
//for j := i + 1; j < sliceLength; j++ {
//fmt.Println(arrayA[j])
//}

//} else {
//fmt.Printf("%d > %d\n", arrayA[i], min)
//}
//}

////fmt.Println(min)
//fmt.Println(arrayA)

//}

func main() {
	arrayA := []int{3, 1, 5, 2, 4}

	min := arrayA[0]
	sliceLength := len(arrayA)

	for i := 1; i < sliceLength; i++ {
		for j := i + 1; j < sliceLength; j++ {
			fmt.Printf("Loop count: %d\n", i)
			if arrayA[i] < min {
				fmt.Printf("Is %d < %d?\n", arrayA[i], min)
				fmt.Printf("yes, min: %d\n", arrayA[i])
				min = i

				//fmt.Println(arrayA[j])
				//}
			}
		}

	}

	fmt.Println(arrayA[min])
	//Why store the min, but just store the index of it.
	//fmt.Println(arrayA)

	//for...

}

//func selectionSort() {

//arrayA := []int{5, 4, 3, 2, 1}
//sliceLength := len(arrayA)

//for i := 0; i < sliceLength; i++ {
//min := i
//for j := i + 1; j < sliceLength; j++ {
//if arrayA[min] > arrayA[j] {
//min = j
//}
//}

//helpers.Swap(&arrayA[i], &arrayA[min])
//}
//}

//iteration 1
//array = 5,4,3,2,1
//i = 0 , min = 0, j = 1
//5 > 4 => true => min = 1

//iteration 2
//i = 0, min = 1 , j = 2
//4 > 3 => true => min = 2

//iteration 3
//3 > 2 => true => min = 3

//iteration 4
//2 > 1 => true => min = 4

//helpers.swap ( arrayA[0], arrayA[4]) => array = 1,4,3,2,5

//iteration 5
//i = 1, min = 1, j = 2
//4 > 3 => true => min = 2

//iteration 6
//i = 1, min 2 , j

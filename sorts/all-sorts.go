package main

import (
	"fmt"

	"github.com/gitflash/go-helpers"
	sort "github.com/gitflash/go-sorts"
)

func main() {
	// Seed Array
	unsrotedArray := helpers.SeedIntSlice(5, 100)

	/* sorts */

	// Bubble sort
	fmt.Println("Buuble Sort -------------------")
	fmt.Println("Input Array:  ", unsrotedArray)
	fmt.Println("Output Array: ", sort.Bubble(unsrotedArray))
	fmt.Println("-------------------------------")

	// Selection sort
	helpers.Shuffle(unsrotedArray) // reshuffle array
	fmt.Println()
	fmt.Println("Select Sort -------------------")
	fmt.Println("Input Array:  ", unsrotedArray)
	fmt.Println("Output Array: ", sort.Selection(unsrotedArray))
	fmt.Println("-------------------------------")

	// Insertion sort
	helpers.Shuffle(unsrotedArray) // reshuffle array
	fmt.Println()
	fmt.Println("Insertion Sort ----------------")
	fmt.Println("Input Array:  ", unsrotedArray)
	fmt.Println("Output Array: ", sort.Insertion(unsrotedArray))
	fmt.Println("-------------------------------")

	// Merge Sort
	helpers.Shuffle(unsrotedArray) // reshuffle array
	fmt.Println()
	fmt.Println("Merge Sort ----------------")
	fmt.Println("Input Array:  ", unsrotedArray)
	fmt.Println("Output Array: ", sort.Merge(unsrotedArray, 0, len(unsrotedArray)-1))
	fmt.Println("-------------------------------")

	// Quick Sort
	helpers.Shuffle(unsrotedArray) // reshuffle array
	fmt.Println()
	fmt.Println("Quick Sort ----------------")
	fmt.Println("Input Array:  ", unsrotedArray)
	fmt.Println("Output Array: ", sort.Quick(unsrotedArray, 0, len(unsrotedArray)-1))
	fmt.Println("-------------------------------")
}

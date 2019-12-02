arrayA = [5, 4, 2, 1]
arrayB = []


//5
min := arrayA[0]

//if 4 < 5 true => min = 4

for i := 1; i < len(arrayA); i++ {
	if array[i] < min 
	  min = array[i]
}

min = 1

arrayB[0] = min

for i:=0 ; i < len(arrayA); i++ {
	min := i
	for j:= i+1; j < len(arrayA); j++ {
		if (arrayA[j] < arrayA[i]) {
			min = j
		}
	}

	helper.Swap(arrayA[i], arrayA[min])
}

// min 0 ( value 5 )
// min 1 ( value 4 )
// min 3 ( value 1 )
// helpers.Swap (array[0], array[3])
// [1,4,2,5]
// arrayA[1] arrayA[2], arrayA[1] arrayA[3]
// [1,2,4,5]

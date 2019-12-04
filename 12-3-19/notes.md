# Insertion sort

Insertion Sort works as follows:

The first step involves the comparison of the element in question with its adjacent element.

[]{1,3,2,4,5}

1 and 3


And if at every comparison reveals that the element in question can be inserted at a particular position, then space is created for it by shifting the other elements one position to the right and inserting the element at the suitable position.


The above procedure is repeated until all the element in the array is at their apt position.
Let us now understand working with the following example:

Consider the following array: 25, 17, 31, 13, 2

First Iteration: Compare 25 with 17. The comparison shows 17< 25. Hence swap 17 and 25.

The array now looks like:

17, 25, 31, 13, 2


Second Iteration: Begin with the second element (25), but it was already swapped on for the correct position, so we move ahead to the next element.

Now hold on to the third element (31) and compare with the ones preceding it.

Since 31> 25, no swapping takes place.

Also, 31> 17, no swapping takes place and 31 remains at its position.

The array after the Second iteration looks like:

17, 25, 31, 13, 2

Third Iteration: Start the following Iteration with the fourth element (13), and compare it with its preceding elements.

Since 13< 31, we swap the two.

Array now becomes: 17, 25, 13, 31, 2.

But there still exist elements that we haven’t yet compared with 13. Now the comparison takes place between 25 and 13. Since, 13 < 25, we swap the two.

The array becomes 17, 13, 25, 31, 2.

The last comparison for the iteration is now between 17 and 13. Since 13 < 17, we swap the two.

The array now becomes 13, 17, 25, 31, 2.

array = 50 , 30 , 40 , 20, 10
1 st iteration 
30 < 50 => swap  => 30, 50 , 40 , 20 , 10

2 nd iteration 
40 < 50 => swap => 30 , 40 , 50, 20 , 10 
40 < 30 => false => 30, 40 , 50 , 20 , 10 

3 rd iteration 
20 < 50 => swap => 30, 40, 20 , 50 , 10
20 < 40 => swap => 30 , 20 , 40 , 50 , 10 
20 < 30 -> swap => 20, 30 , 40 , 50 , 10 

4 th 




arrayA := []int{50, 30, 40, 20, 10}

//Pseudo code...

for i := 0; i < len(arrayA); i++ {
	for j := 0...
	//if 1 < 0?
	if array[1] < array[0] {
		//50
		tmp = array[0]
		array[0] = array[1]
		array[1] = tmp
	}
}




























# Bubble sort

Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
Example:
First Pass:

```
( 5 1 4 2 8 ) –> ( 1 5 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.
( 1 5 4 2 8 ) –>  ( 1 4 5 2 8 ), Swap since 5 > 4
( 1 4 5 2 8 ) –>  ( 1 4 2 5 8 ), Swap since 5 > 2
( 1 4 2 5 8 ) –> ( 1 4 2 5 8 ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.
```
Second Pass:
```
( 1 4 2 5 8 ) –> ( 1 4 2 5 8 )
( 1 4 2 5 8 ) –> ( 1 2 4 5 8 ), Swap since 4 > 2
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
```

Now, the array is already sorted, but our algorithm does not know if it is completed. The algorithm needs one whole pass without any swap to know it is sorted.
Third Pass:
```
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
( 1 2 4 5 8 ) –> ( 1 2 4 5 8 )
```

1. Pick a target reference. (Pivot)

```
list := (5, 4, 3, 1, 2)
```

Pivot = 4 dk - unsure how the pivot is picked...

- input, ouput

Bubble sort - input is array, output is array.
- Descending sort or ascending sort. We will choose ascending. (Be sure to clarify this and not just start coding...)

[3,2,1] => [1,2,3]

[ 
	0 => 3
	1 => 2
	2 => 1
]

startloop i = 0 end = N(2)
	if (arr[i] > arr[i+1])
		// swap arr[i] with arr[i+1]
endloop

[3,2,1]

1st iteration [2,3,1]
2nd iteration [2,1,3]

bad solution


```
for i = 0 ; i <= arr.length - 1; i++
	for j = i + 1; j <= arr.length - 1; j++ 	
		//i = 0 + 1  if lt 3 - 1 = 2
		if (arr[i] > arr[j]) // swap
			//arr[i] = 0 or value... Not correct
			//3 > 
	endfor
endfor
```

array = [a,b,c,d]

```
array = [
	0 => a,
	1 => b,
	2 => c,
	3 => d
]
```

```
for var i = 0
	exit condition: if i lt or equal (length of array - 1)
    step: i = i+1
	 
   echo array[i]
	
endfor
```

result : abcd

for 


exit (when does the exit happen?)

What does compiler do? 

1. declares a variable called "i"
2. initiate the variable i the value of 0
3. execute the exit condition if condition true continue else exit
4. if continue : execute the increment ( step )
5. if continue : go back to 3

i = 0 ; i < 3 = yes , echo array[0], i = i + 1, if i < 3 = yes, echo array[1], i = i + 1 
exit 

```
0th iteration [4,3,2,1] 

arr[0] = 4 , arr[1] = 3 , arr[2] = 2, arr[3] = 1

1st iteration i = 0, j = 1    => [3,4,2,1]
2nd iteration i = 0, j = 2 	=> [2,4,3,1]
3rd iteration i = 0, j = 3		=> [1,4,3,2]
4th iteration i = 0, j = 4    => fail 
5th iteration i = 1, j = 2    => [1,3,4,2]
6th iteration i = 1, j = 3		=> [1,2,4,3]
```

#implement for while loop

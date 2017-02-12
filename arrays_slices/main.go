package main

import (
	"fmt"
)

func main() {
	//fixed arrays
	var x [5]int
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x, arr)
	//slices
	// zero elements
	var y []int
	fmt.Println("zero elements", y)
	// index 5 excluded
	y = arr[0:5]
	fmt.Println("exlude index", y)
	//creates a new slice of 5 lenght
	y = make([]int, 5)
	fmt.Println(y, len(y))
	//creates a new slice of 5 capacity
	y = make([]int, 5, 10)
	fmt.Println(y, len(y))
	y = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(y, len(y))
	//add elements to slice
	slice1 := append(y, 6, 8)
	fmt.Println("slice1:", slice1)
	slice2 := make([]int, 2)
	//only copies the elements that fixed into the slice1
	copied := copy(slice2, slice1)
	fmt.Println(copied, slice2)

}

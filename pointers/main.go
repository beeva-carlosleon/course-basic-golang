package main

import "fmt"

func main() {
	var myPointer *int
	fmt.Printf("'%v'\n", myPointer)
	//fmt.Printf("'%v'\n", *myPointer)
	myPointer = new(int)
	fmt.Printf("Reference '%v'\n", myPointer)
	fmt.Printf("Value '%v'\n", *myPointer)
	myVar := 1
	myPointer = &myVar
	fmt.Printf("Reference '%v'\n", myPointer)
	fmt.Printf("Value '%v'\n", *myPointer)
}

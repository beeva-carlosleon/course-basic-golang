package main

import "fmt"

func main() {
	//loops
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	//range loops
	myArray := []string{"one", "two"}
	for index, value := range myArray {
		fmt.Println(index, value)
	}
	//while statements
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//infinite loop
	/*for {
		fmt.Println(i)
	}*/
	//if statements
	if i == 10 {
		fmt.Println("Enter")
	} else if i < 1 {
		fmt.Println("Enter?")
	} else {
		fmt.Println("Never Enter")
	}
	//case statements
	switch i {
	case 1, 2:
		fmt.Println("Never enter")
	case 3:
		fmt.Println("Never enter")
	default:
		fmt.Println("Enter")
	}
}

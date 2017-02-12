package main

import (
	"errors"
	"fmt"
)

func myProcedure() {
	fmt.Println("Enters procedure")
}

func sum(parameter1 int, parameter2, parameter3 int) int {
	return parameter1 + parameter2 + parameter3
}

func sum2(parameter1 int, parameter2, parameter3 int) (sumValue int) {
	sumValue = parameter1 + parameter2 + parameter3
	return
}

func sum3(parameter1 int, parameter2, parameter3 int) (int, error) {
	if parameter1 == 1 {
		return 0, errors.New("Invalid parameter")
	}
	return parameter1 + parameter2 + parameter3, nil
}

func sum4(parameters ...int) (result int) {
	for _, value := range parameters {
		result += value
	}
	return
}

func Doubles() func() int {
	i := 1
	return func() (ret int) {
		i *= 2
		ret = i
		return ret
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	myProcedure()
	//defer myProcedure()
	defer func() {
		fmt.Println("Catch error:", recover())
	}()
	fmt.Printf("Sum: '%d'\n", sum(1, 2, 3))
	fmt.Printf("Sum2: '%d'\n", sum2(4, 5, 6))
	result, status := sum3(1, 5, 6)
	fmt.Printf("Sum3 KO: '%d', Error '%v'\n", result, status)
	result, status = sum3(4, 5, 6)
	fmt.Printf("Sum3 OK: '%d', '%v'\n", result, status)
	fmt.Printf("Variadic sum: '%d'\n", sum4(1, 2, 3, 4, 5, 6))
	doubleRef := Doubles()
	fmt.Printf("Closure double: '%d'\n", doubleRef())
	fmt.Printf("Closure double: '%d'\n", doubleRef())
	fmt.Printf("Recursion factorial: '%d'\n", factorial(3))
	panic("Raise an error")
}

package main

import "fmt"

const (
	GlobalConst = "global const"
	localConst  = "local const"
)

var (
	localVariable  int
	GlobalVariable string = "global"
)

func main() {
	//constants
	fmt.Printf("Local Contant '%s'\n", localConst)
	fmt.Printf("Global Contant '%s'\n", GlobalConst)
	//local module variable
	fmt.Printf("Local variable int default Value '%d'\n", localVariable)
	//global module variable
	fmt.Printf("Global variable default Value '%s'\n", GlobalVariable)
	//define inner variable with type
	var myString string
	fmt.Printf("string Default Value '%s'\n", myString)
	myString = "New value"
	fmt.Printf("Value '%s'\n", myString)
	//define variable without type
	myInt := 1
	fmt.Printf("Int Value '%d'\n", myInt)
	//casting
	fmt.Printf("Casted Value to float '%f'\n", float32(myInt))
	//type assertions
	var myAssertion interface{} = 1
	fmt.Printf("myAssertion Value '%d'\n", myAssertion)
	myVar, ok := myAssertion.(int64)
	fmt.Printf("myVar Value '%d', assertion status '%t'\n", myVar, ok)
	//myVar=myAssertion.(int64)
	//fmt.Printf("myVar Value '%d', assertion status %t\n", myVar, ok)
	switch value := myAssertion.(type) {
	case int:
		fmt.Printf("Value '%d'\n", value)
	case string:
		fmt.Print("never enters\n")
	}
}

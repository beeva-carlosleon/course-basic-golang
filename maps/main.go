package main

import (
	"fmt"
)

func main() {
	var y map[string]int = make(map[string]int) //must  be initialized
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"], y)
	delete(x, "key")
	fmt.Println(x, y)
	element, ok := x["Un"]
	fmt.Println(element, ok)
}

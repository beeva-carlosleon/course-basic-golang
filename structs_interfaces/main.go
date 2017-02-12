package main

import (
	"fmt"
	"github.com/curso/structs_interfaces/structs"
)

func main() {
	circle := structs.NewCircle()
	circle.SetFields(0, 0, 10)
	fmt.Println("Area", circle.Area())

	circle2 := &structs.Circle{
		R: 10,
		X: 0,
		Y: 0,
	}
	fmt.Println("Area2", circle2.Area())
	circle2.SetRadius(20)
	fmt.Println("Area2", circle2.Area())
	circle = structs.NewRectangle()
	circle.SetFields(0, 10, 0, 10)
	fmt.Println("Length", circle.Length())
	circle = structs.NewRectangle2()
	circle.SetFields(0, 10, 0, 10)
	fmt.Println("Length", circle.Length())
}

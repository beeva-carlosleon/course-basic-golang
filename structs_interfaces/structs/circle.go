package structs

import (
	"github.com/course-basic-golang/structs_interfaces/interfaces"
	"math"
)

type Circle struct {
	X float64
	Y float64
	R float64
}

func NewCircle() interfaces.Shape {
	return new(Circle)
}

func (c *Circle) SetRadius(r float64) {
	c.R = r
}

func (c *Circle) SetFields(fields ...float64) {
	c.X = fields[0]
	c.Y = fields[1]
	c.R = fields[2]
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Length() float64 {
	return 2 * math.Pi * c.R
}

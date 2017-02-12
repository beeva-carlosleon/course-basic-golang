package structs

import (
	"math"

	"github.com/curso/structs_interfaces/interfaces"
)

type Rectangle struct {
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func NewRectangle() interfaces.Shape {
	return new(Rectangle)
}

func (r *Rectangle) SetFields(fields ...float64) {
	r.X1 = fields[0]
	r.X2 = fields[1]
	r.Y1 = fields[2]
	r.Y2 = fields[3]
}

func (r Rectangle) Area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l * w
}

func (r Rectangle) Length() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l*2 + w*2
}

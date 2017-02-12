package structs

import "github.com/course-basic-golang/structs_interfaces/interfaces"

type Point struct {
	X, Y float64
}

type Rectangle2 struct {
	P1 Point
	P2 Point
}

func NewRectangle2() interfaces.Shape {
	return new(Rectangle2)
}

func (r *Rectangle2) SetFields(fields ...float64) {
	r.P1.X = fields[0]
	r.P2.X = fields[1]
	r.P1.Y = fields[2]
	r.P2.Y = fields[3]
}

func (r Rectangle2) Area() float64 {
	l := distance(r.P1.X, r.P1.Y, r.P1.X, r.P2.X)
	w := distance(r.P1.X, r.P1.Y, r.P2.X, r.P1.Y)
	return l * w
}

func (r Rectangle2) Length() float64 {
	l := distance(r.P1.X, r.P1.Y, r.P1.X, r.P2.Y)
	w := distance(r.P1.X, r.P1.Y, r.P2.X, r.P1.Y)
	return l*2 + w*2
}

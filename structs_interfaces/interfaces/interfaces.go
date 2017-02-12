package interfaces

type Fields interface {
	SetFields(fields ...float64)
}
type Shape interface {
	Length() float64
	Area() float64
	Fields
}

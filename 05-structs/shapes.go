package structs

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Triangle struct {
	Length float64
	Height float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (t Triangle) Area() float64 {
	return t.Length * t.Height * 0.5
}

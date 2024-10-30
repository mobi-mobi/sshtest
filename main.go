package main

import "fmt"

type shape interface {
	Area() float64
}
type square struct {
	a float64
}

type rectangle struct {
	a float64
	b float64
}

func (s square) Area() float64 {
	return s.a * s.a
}

func (r rectangle) Area() float64 {
	return r.a * r.b
}

func main() {
	stvorec := square{4}
	obdlznik := rectangle{4, 3}

	shapes := []shape{stvorec, obdlznik}
	fmt.Println(shapes[1].Area())
}

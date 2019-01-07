package main

import (
	"fmt"

	"./shape"
)

func main() {
	c := new(shape.Circle)
	r := new(shape.Rectangle)
	s := new(shape.Square)
	c.SetRadius(2)
	r.SetLength(2)
	r.SetWidth(2)
	s.SetLength(2)
	fmt.Println("area of circle: ", c.Area())
	fmt.Println("circumference of circle: ", c.Circumference())
	fmt.Println("area of rectangle: ", r.Area())
	fmt.Println("perimeter of rectangle: ", r.Perimeter())
	fmt.Println("area of square:  ", s.Area())
	fmt.Println("perimeter of square: ", s.Perimeter())

}

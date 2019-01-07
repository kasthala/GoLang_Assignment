package shape

import "math"

//Circle xyz
type Circle struct {
	radius float64
}

//Area xyz
func (c Circle) Area() float64 {

	return math.Pi * c.Radius() * c.Radius()
}

//Circumference xyz
func (c Circle) Circumference() float64 {

	return 2 * math.Pi * c.Radius()
}

//SetRadius xyz
func (c *Circle) SetRadius(r float64) {
	c.radius = r
}

//Radius xyz
func (c Circle) Radius() float64 {
	return c.radius
}

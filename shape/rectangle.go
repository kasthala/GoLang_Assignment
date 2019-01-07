package shape

//Rectangle  xyz
type Rectangle struct {
	width, length float64
}

//Area xyz
func (r Rectangle) Area() float64 {

	return r.length * r.width
}

//Perimeter xyz
func (r Rectangle) Perimeter() float64 {

	return 2 * (r.length + r.width)
}

//SetWidth xyz
func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

//Width xyz
func (r Rectangle) Width() float64 {
	return r.width
}

//SetLength xyz
func (r *Rectangle) SetLength(l float64) {
	r.length = l
}

//Length xyz
func (r Rectangle) Length() float64 {
	return r.length
}

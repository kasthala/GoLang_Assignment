package shape

//Square xyz
type Square struct {
	length float64
}

//Area xyz
func (s Square) Area() float64 {

	return s.length * s.length
}

//Perimeter xyz
func (s Square) Perimeter() float64 {

	return 4 * s.length
}

//SetLength xyz
func (s *Square) SetLength(l float64) {
	s.length = l
}

//Length xyz
func (s Square) Length() float64 {
	return s.length
}

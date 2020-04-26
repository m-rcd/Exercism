// package triange determines the type of a triangle
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides return the type of a triangle.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case NotATriangle(a, b, c):
		return NaT
	case a == b && b == c && a != 0:
		return Equ
	case (a == b && b != c) || (a == c && c != b) || (b == c && c != a):
		return Iso
	case a != b && b != c:
		return Sca
	}
	return NaT
}

func NotATriangle(a, b, c float64) bool {
	return DegenerateTriangle(a, b, c) || SidesNotANumber(a, b, c) || SidesNotInf(a, b, c)
}

func SidesNotANumber(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
}

func SidesNotInf(a, b, c float64) bool {
	return math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}

func DegenerateTriangle(a, b, c float64) bool {
	return (a+b < c) || (a+c < b) || (b+c < a)
}

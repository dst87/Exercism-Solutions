// Package triangle offers methods for assessing whether certain
// meet the requirements to be considered triangles.
package triangle

import (
	"sort"
	"math"
)

type Kind int

const (
	NaT Kind = 0 // not a triangle
	Equ Kind = 1 // equilateral
	Iso Kind = 2 // isosceles
	Sca Kind = 3 // scalene
)

// KindFromSides takes the lengths of the sides of a proposed triangle and
// returns the type (or that the input could not be a triangle).
func KindFromSides(a, b, c float64) Kind {

	sides := []float64{a, b, c}

	for _, side := range sides {
		if side == 0 || math.IsInf(side, 0) || math.IsNaN(side) {
			// Inputs where any side is 0, negative (Not a Number),
			// or infinite are not triangles.
			return NaT
		}
	}
		
	if (a == b) && (b == c) {
		// Triangles where all sides are equal
		return Equ
	}

	sort.Float64s(sides)

	if sides[0]+sides[1] < sides[2] {
		return NaT
	}

	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}

	return Sca
}

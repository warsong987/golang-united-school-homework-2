package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return CalcTiangleSquare(sideLen)

	case SidesSquare:
		return CalcSquareSquare(sideLen)

	case SidesCircle:
		return CalcSquareOval(sideLen)
	}
	return -1
}

func CalcTiangleSquare(sideLen float64) float64 {
	return (math.Sqrt(3) / 4) * (sideLen * sideLen)
}

func CalcSquareSquare(sideLen float64) float64 {
	return sideLen * sideLen
}

func CalcSquareOval(radius float64) float64 {
	return math.Pi * radius * radius
}

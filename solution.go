package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type sides int

const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pow(sideLen, 2) * math.Pi * 0.5
	case SidesTriangle:
		return (math.Sqrt(3) * sideLen) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	}

	return 0
}

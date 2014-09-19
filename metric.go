package simpleline

import (
	"math"
)

type Metric func(Point, Point) float64

// Euclidean implements the euclidean distance metric
func Euclidean(p1, p2 Point) float64 {
	vec1 := p1.Vector()
	vec2 := p2.Vector()

	if len(vec1) != len(vec2) {
		panic("Vectors are different lengths")
	}

	sum := float64(0)

	iDiff := float64(0)

	for i := 0; i < len(vec1); i++ {
		iDiff = vec1[i] - vec2[i]
		sum += iDiff * iDiff
	}

	return math.Sqrt(sum)
}

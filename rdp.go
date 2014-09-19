// Package simpleline simplifies N dimensional lines using the Ramer–Douglas–Peucker algorithm.
// This package is not the highest performing implementation of the RDP algorithm.
package simpleline

import (
	"errors"
	"math"
)

func RDP(points []Point, epsilon float64, d Metric, simplifyFirst bool) (results []Point, err error) {
	defer func() {
		if r := recover(); r != nil {
			if str, ok := r.(string); ok {
				err = errors.New(str)
			}
			err = r.(error)
		}
	}()

	if len(points) <= 1 {
		return points, nil
	}

	if simplifyFirst {
		simplifiedPoints := simplifyRadialDist(points, epsilon, d)
		results = rdp(simplifiedPoints, epsilon, d)
	} else {
		results = rdp(points, epsilon, d)
	}

	return
}

func rdp(points []Point, epsilon float64, d Metric) []Point {
	if len(points) <= 1 {
		return points
	}

	stack := []int{}
	length := len(points)
	markers := make([]bool, length)

	dMax := float64(0)
	maxIndex := -1

	first := 0
	last := len(points) - 1
	markers[first], markers[last] = true, true

	stack = append(stack, first, last)

	for stackLen := len(stack); stackLen > 0; stackLen = len(stack) {
		last = stack[stackLen-1]
		first = stack[stackLen-2]
		stack = stack[:stackLen-2]

		dMax = 0

		for i := first + 1; i < last; i++ {
			dist := shortestDistanceToSegment(points[first], points[last], points[i], d)

			if dist > dMax {
				maxIndex = i
				dMax = dist
			}
		}

		if dMax > epsilon {
			markers[maxIndex] = true
			stack = append(stack, first, maxIndex, maxIndex, last)
		}

	}

	results := []Point{}

	for i, include := range markers {
		if include {
			results = append(results, points[i])
		}
	}

	return results
}

func shortestDistanceToSegment(a1, a2 Point, point Point, d Metric) float64 {
	slope := a1.Subtract(a2)

	slopeLength := d(slope.Zero(), slope)

	if slopeLength == 0 {
		return d(a1, point)
	}

	unitSlope := slope.Scale(float64(1) / slopeLength)

	toPoint := point.Subtract(a1)

	shortestLine := toPoint.Subtract(unitSlope.Scale(dot(toPoint, unitSlope)))

	toLine := d(shortestLine.Zero(), shortestLine)

	toA1 := d(a1, point)
	toA2 := d(a2, point)

	veryShortest := math.Min(toLine, math.Min(toA1, toA2))

	return veryShortest
}

// basic distance-based simplification
func simplifyRadialDist(points []Point, epsilon float64, d Metric) []Point {

	prevPoint := points[0]

	newPoints := []Point{prevPoint}

	var point Point

	for i := 1; i < len(points); i++ {
		point = points[i]

		if d(point, prevPoint) > epsilon {
			newPoints = append(newPoints, point)
			prevPoint = point
		}
	}

	if point != nil && prevPoint != point {
		newPoints = append(newPoints, point)
	}

	return newPoints
}

func dot(p1, p2 Point) float64 {
	vec1 := p1.Vector()
	vec2 := p2.Vector()

	if len(vec1) != len(vec2) {
		panic("Vectors are different lengths")
	}

	sum := float64(0)
	for i, a1 := range vec1 {
		sum += a1 * vec2[i]
	}
	return sum
}

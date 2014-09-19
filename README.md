# simpleline

Simpleline simplifies N dimensional lines using the Ramer–Douglas–Peucker algorithm.

It is a port of https://github.com/mourner/simplify-js by Vladimir Agafonkin, extended to support N dimensions and custom metrics (distance functions).

## Usage

To use simpleline, you provide a slice of Point objects.

Point is an interface defined by:

	type Point interface {
		Vector() []float64
		Scale(float64) Point
		Subtract(Point) Point
		Zero() Point
	}

A three dimensional point, Point3d, is provided. You can use use it in 2D if you want.

You simplify your list of points by calling 
	
	points := []Point{...}

	results, err := simpleline.RDP(points, 5, simpleline.Euclidean, true)
	if err != nil {
		// handle error
	}

simpleline.RDP is called like:

	simpleline.RDP(points, epsilon, metric function, simplify first?)

## Custom metric

The Euclidean metric is provided. You can implement your own function that satisfies the Metric interface.

	type Metric func(Point, Point) float64

## Caveats

You could get better performance. For example, this implementation uses vector math and calculates the actual distance rather than its square.

Be careful with floats near MaxFloat64.
package simpleline

type Point interface {
	Vector() []float64
	Scale(float64) Point
	Subtract(Point) Point
	Zero() Point
}

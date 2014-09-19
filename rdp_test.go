package simpleline

import (
	"testing"
)

// test data from https://github.com/mourner/simplify-js/blob/master/test.js

var testPoints []Point = []Point{
	&Point3d{X: 224.55, Y: 250.15, Z: 0}, &Point3d{X: 226.91, Y: 244.19, Z: 0}, &Point3d{X: 233.31, Y: 241.45, Z: 0}, &Point3d{X: 234.98, Y: 236.06, Z: 0},
	&Point3d{X: 244.21, Y: 232.76, Z: 0}, &Point3d{X: 262.59, Y: 215.31, Z: 0}, &Point3d{X: 267.76, Y: 213.81, Z: 0}, &Point3d{X: 273.57, Y: 201.84, Z: 0},
	&Point3d{X: 273.12, Y: 192.16, Z: 0}, &Point3d{X: 277.62, Y: 189.03, Z: 0}, &Point3d{X: 280.36, Y: 181.41, Z: 0}, &Point3d{X: 286.51, Y: 177.74, Z: 0},
	&Point3d{X: 292.41, Y: 159.37, Z: 0}, &Point3d{X: 296.91, Y: 155.64, Z: 0}, &Point3d{X: 314.95, Y: 151.37, Z: 0}, &Point3d{X: 319.75, Y: 145.16, Z: 0},
	&Point3d{X: 330.33, Y: 137.57, Z: 0}, &Point3d{X: 341.48, Y: 139.96, Z: 0}, &Point3d{X: 369.98, Y: 137.89, Z: 0}, &Point3d{X: 387.39, Y: 142.51, Z: 0},
	&Point3d{X: 391.28, Y: 139.39, Z: 0}, &Point3d{X: 409.52, Y: 141.14, Z: 0}, &Point3d{X: 414.82, Y: 139.75, Z: 0}, &Point3d{X: 427.72, Y: 127.30, Z: 0},
	&Point3d{X: 439.60, Y: 119.74, Z: 0}, &Point3d{X: 474.93, Y: 107.87, Z: 0}, &Point3d{X: 486.51, Y: 106.75, Z: 0}, &Point3d{X: 489.20, Y: 109.45, Z: 0},
	&Point3d{X: 493.79, Y: 108.63, Z: 0}, &Point3d{X: 504.74, Y: 119.66, Z: 0}, &Point3d{X: 512.96, Y: 122.35, Z: 0}, &Point3d{X: 518.63, Y: 120.89, Z: 0},
	&Point3d{X: 524.09, Y: 126.88, Z: 0}, &Point3d{X: 529.57, Y: 127.86, Z: 0}, &Point3d{X: 534.21, Y: 140.93, Z: 0}, &Point3d{X: 539.27, Y: 147.24, Z: 0},
	&Point3d{X: 567.69, Y: 148.91, Z: 0}, &Point3d{X: 575.25, Y: 157.26, Z: 0}, &Point3d{X: 580.62, Y: 158.15, Z: 0}, &Point3d{X: 601.53, Y: 156.85, Z: 0},
	&Point3d{X: 617.74, Y: 159.86, Z: 0}, &Point3d{X: 622.00, Y: 167.04, Z: 0}, &Point3d{X: 629.55, Y: 194.60, Z: 0}, &Point3d{X: 638.90, Y: 195.61, Z: 0},
	&Point3d{X: 641.26, Y: 200.81, Z: 0}, &Point3d{X: 651.77, Y: 204.56, Z: 0}, &Point3d{X: 671.55, Y: 222.55, Z: 0}, &Point3d{X: 683.68, Y: 217.45, Z: 0},
	&Point3d{X: 695.25, Y: 219.15, Z: 0}, &Point3d{X: 700.64, Y: 217.98, Z: 0}, &Point3d{X: 703.12, Y: 214.36, Z: 0}, &Point3d{X: 712.26, Y: 215.87, Z: 0},
	&Point3d{X: 721.49, Y: 212.81, Z: 0}, &Point3d{X: 727.81, Y: 213.36, Z: 0}, &Point3d{X: 729.98, Y: 208.73, Z: 0}, &Point3d{X: 735.32, Y: 208.20, Z: 0},
	&Point3d{X: 739.94, Y: 204.77, Z: 0}, &Point3d{X: 769.98, Y: 208.42, Z: 0}, &Point3d{X: 779.60, Y: 216.87, Z: 0}, &Point3d{X: 784.20, Y: 218.16, Z: 0},
	&Point3d{X: 800.24, Y: 214.62, Z: 0}, &Point3d{X: 810.53, Y: 219.73, Z: 0}, &Point3d{X: 817.19, Y: 226.82, Z: 0}, &Point3d{X: 820.77, Y: 236.17, Z: 0},
	&Point3d{X: 827.23, Y: 236.16, Z: 0}, &Point3d{X: 829.89, Y: 239.89, Z: 0}, &Point3d{X: 851.00, Y: 248.94, Z: 0}, &Point3d{X: 859.88, Y: 255.49, Z: 0},
	&Point3d{X: 865.21, Y: 268.53, Z: 0}, &Point3d{X: 857.95, Y: 280.30, Z: 0}, &Point3d{X: 865.48, Y: 291.45, Z: 0}, &Point3d{X: 866.81, Y: 298.66, Z: 0},
	&Point3d{X: 864.68, Y: 302.71, Z: 0}, &Point3d{X: 867.79, Y: 306.17, Z: 0}, &Point3d{X: 859.87, Y: 311.37, Z: 0}, &Point3d{X: 860.08, Y: 314.35, Z: 0},
	&Point3d{X: 858.29, Y: 314.94, Z: 0}, &Point3d{X: 858.10, Y: 327.60, Z: 0}, &Point3d{X: 854.54, Y: 335.40, Z: 0}, &Point3d{X: 860.92, Y: 343.00, Z: 0},
	&Point3d{X: 856.43, Y: 350.15, Z: 0}, &Point3d{X: 851.42, Y: 352.96, Z: 0}, &Point3d{X: 849.84, Y: 359.59, Z: 0}, &Point3d{X: 854.56, Y: 365.53, Z: 0},
	&Point3d{X: 849.74, Y: 370.38, Z: 0}, &Point3d{X: 844.09, Y: 371.89, Z: 0}, &Point3d{X: 844.75, Y: 380.44, Z: 0}, &Point3d{X: 841.52, Y: 383.67, Z: 0},
	&Point3d{X: 839.57, Y: 390.40, Z: 0}, &Point3d{X: 845.59, Y: 399.05, Z: 0}, &Point3d{X: 848.40, Y: 407.55, Z: 0}, &Point3d{X: 843.71, Y: 411.30, Z: 0},
	&Point3d{X: 844.09, Y: 419.88, Z: 0}, &Point3d{X: 839.51, Y: 432.76, Z: 0}, &Point3d{X: 841.33, Y: 441.04, Z: 0}, &Point3d{X: 847.62, Y: 449.22, Z: 0},
	&Point3d{X: 847.16, Y: 458.44, Z: 0}, &Point3d{X: 851.38, Y: 462.79, Z: 0}, &Point3d{X: 853.97, Y: 471.15, Z: 0}, &Point3d{X: 866.36, Y: 480.77, Z: 0},
}

var testPointsSimplified []Point = []Point{
	&Point3d{X: 224.55, Y: 250.15, Z: 0}, &Point3d{X: 267.76, Y: 213.81, Z: 0}, &Point3d{X: 296.91, Y: 155.64, Z: 0}, &Point3d{X: 330.33, Y: 137.57, Z: 0},
	&Point3d{X: 409.52, Y: 141.14, Z: 0}, &Point3d{X: 439.60, Y: 119.74, Z: 0}, &Point3d{X: 486.51, Y: 106.75, Z: 0}, &Point3d{X: 529.57, Y: 127.86, Z: 0},
	&Point3d{X: 539.27, Y: 147.24, Z: 0}, &Point3d{X: 617.74, Y: 159.86, Z: 0}, &Point3d{X: 629.55, Y: 194.60, Z: 0}, &Point3d{X: 671.55, Y: 222.55, Z: 0},
	&Point3d{X: 727.81, Y: 213.36, Z: 0}, &Point3d{X: 739.94, Y: 204.77, Z: 0}, &Point3d{X: 769.98, Y: 208.42, Z: 0}, &Point3d{X: 779.60, Y: 216.87, Z: 0},
	&Point3d{X: 800.24, Y: 214.62, Z: 0}, &Point3d{X: 820.77, Y: 236.17, Z: 0}, &Point3d{X: 859.88, Y: 255.49, Z: 0}, &Point3d{X: 865.21, Y: 268.53, Z: 0},
	&Point3d{X: 857.95, Y: 280.30, Z: 0}, &Point3d{X: 867.79, Y: 306.17, Z: 0}, &Point3d{X: 859.87, Y: 311.37, Z: 0}, &Point3d{X: 854.54, Y: 335.40, Z: 0},
	&Point3d{X: 860.92, Y: 343.00, Z: 0}, &Point3d{X: 849.84, Y: 359.59, Z: 0}, &Point3d{X: 854.56, Y: 365.53, Z: 0}, &Point3d{X: 844.09, Y: 371.89, Z: 0},
	&Point3d{X: 839.57, Y: 390.40, Z: 0}, &Point3d{X: 848.40, Y: 407.55, Z: 0}, &Point3d{X: 839.51, Y: 432.76, Z: 0}, &Point3d{X: 853.97, Y: 471.15, Z: 0},
	&Point3d{X: 866.36, Y: 480.77, Z: 0},
}

func TestRDPin2DOnePoint(t *testing.T) {
	onePoint := []Point{&Point3d{X: 1, Y: 2}}
	oneResult, err := RDP(onePoint, 5, Euclidean, true)
	if err != nil {
		t.Fail()
	}
	if len(oneResult) != 1 {
		t.Log("Returned not one point")
		t.Fail()
	}
}

func TestRDPin2DZeroPoints(t *testing.T) {
	noPoints := []Point{}
	noResult, err := RDP(noPoints, 5, Euclidean, true)
	if err != nil {
		t.Fail()
	}
	if len(noResult) != 0 {
		t.Log("Returned zero points")
		t.Fail()
	}
}

func TestRDPin2D(t *testing.T) {
	onePoint := []Point{&Point3d{X: 1, Y: 2}}
	oneResult, err := RDP(onePoint, 5, Euclidean, true)
	if err != nil {
		t.Fail()
	}
	if len(oneResult) != 1 {
		t.Log("Returned not one point")
		t.Fail()
	}

	results, err := RDP(testPoints, 5, Euclidean, true)
	if err != nil {
		t.Fail()
	}

	t.Log("Input point count:", len(testPoints))

	if len(results) != len(testPointsSimplified) {
		t.Log("Results had length", len(results), "but expected length was", len(testPointsSimplified))
		t.Fail()
	}

}

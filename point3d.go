package simpleline

type Point3d struct {
	X, Y, Z float64
}

func (p *Point3d) Vector() []float64 {
	return []float64{p.X, p.Y, p.Z}
}

func (p *Point3d) Scale(val float64) Point {
	np := Point3d{}
	np.X = p.X * val
	np.Y = p.Y * val
	np.Z = p.Z * val
	return &np
}

func (p *Point3d) Subtract(p2 Point) Point {
	p2Cast := p2.(*Point3d)
	np := Point3d{}
	np.X = p.X - p2Cast.X
	np.Y = p.Y - p2Cast.Y
	np.Z = p.Z - p2Cast.Z
	return &np
}

func (p *Point3d) Zero() Point {
	np := Point3d{}
	return &np
}

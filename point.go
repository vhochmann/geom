package geom

// abs returns the absolute value of an int
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Point represents a point in 2 dimensional Euclidean space
type Point struct {
	X, Y int
}

// Pt creates a new point at x, y
func Pt(x, y int) Point {
	return Point{X:x, Y:y}
}

// Add combines points into a new one
func (p Point) Add(other Point) Point {
	return Pt(p.X+other.X, p.Y+other.Y)
}

// Sub returns a new point representing the difference between the two
func (p Point) Sub(other Point) Point {
	return Pt(p.X-other.X, p.Y-other.Y)
}

// Mult multiplies two points and returns a new one with that value
func (p Point) Mult(other Point) Point {
	return Pt(p.X*other.X, p.Y*other.Y)
}

// Div divides two points and returns a new one with that value
func (p Point) Div(other Point) Point {
	return Pt(p.X/other.X, p.Y/other.Y)
}

// Mid returns the midpoint of two Points
func (p Point) Mid(other Point) Point {
	return p.Add(other).Div(Pt(2,2))
}

// Dist returns the Euclidean distance between two points
func (p Point) Dist(other Point) (out int) {
	var res Point
	res = other.Sub(p)
	out = abs(res.X)+abs(res.Y)
	return
}

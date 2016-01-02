package geom

// A Shape contains points
type Shape interface{
	Points() []Point
}

// abs returns the absolute value of a passed int
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

package geom

type Line struct {
	P1, P2 Point
}

func NewLine(p1, p2 Point) Line {
	return Line{P1:p1, P2: p2}
}

// Points is an implementation of Bresenham's Line Algorithm, copied
// directly from roguebasin.com. It returns all points line
func (l Line) Points() (points []Point) {
	x1, y1 := l.P1.X, l.P1.Y
	x2, y2 := l.P2.X, l.P2.Y
 
	isSteep := abs(y2-y1) > abs(x2-x1)
	if isSteep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}
 
	reversed := false
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
		reversed = true
	}
 
	deltaX := x2 - x1
	deltaY := abs(y2 - y1)
	err := deltaX / 2
	y := y1
	var ystep int
 
	if y1 < y2 {
		ystep = 1
	} else {
		ystep = -1
	}
 
	for x := x1; x < x2+1; x++ {
		if isSteep {
			points = append(points, Pt(y, x))
		} else {
			points = append(points, Pt(x, y))
		}
		err -= deltaY
		if err < 0 {
			y += ystep
			err += deltaX
		}
	}
 
	if reversed {
		for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
			points[i], points[j] = points[j], points[i]
		}
	}
 
	return
}

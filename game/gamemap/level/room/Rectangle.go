package room

// Rectangle represent top left and bottom right points of the rectangle room
type Rectangle struct {
	X1, Y1, X2, Y2 int
}

func NewRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{X1: x, Y1: y, X2: x + width, Y2: y + height}
}

func (r *Rectangle) Center() (int, int) {
	return (r.X2 + r.X1) / 2, (r.Y2 + r.Y1) / 2
}

func (r *Rectangle) IsIntersect(other *Rectangle) bool {
	return (r.X1 <= other.X2 && r.X2 >= other.X1) && (r.Y1 <= other.Y1 && r.Y2 >= other.Y1)
}

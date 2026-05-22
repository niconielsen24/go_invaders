package internal

type BoundingBox struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func NewBoundingBox(x, y, width, height int32) BoundingBox {
	return BoundingBox{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (b *BoundingBox) Contains(p Point) bool {
	return p.X >= b.X && p.X <= b.X+b.Width &&
		p.Y >= b.Y && p.Y <= b.Y+b.Height
}

func (b *BoundingBox) Intersects(other BoundingBox) bool {
	return b.X < other.X+other.Width &&
		b.X+b.Width > other.X &&
		b.Y < other.Y+other.Height &&
		b.Y+b.Height > other.Y
}

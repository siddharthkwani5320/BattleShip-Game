package component

type Point struct {
	row int
	col int
}

var ShipPoint = make(map[int][]*Point)

func NewPoint(row, col int) *Point {
	point := &Point{
		row: row,
		col: col,
	}
	return point
}

func (p *Point) GetRow() int {
	return p.row
}

func (p *Point) GetCol() int {
	return p.col
}

package component

type Mark string

const (
	NoMark   Mark = "-"
	HitMark  Mark = "H"
	MissMark Mark = "M"
)

type Cells struct {
	mark Mark
}

func NewCell() *Cells {
	var tempCell = &Cells{
		mark: NoMark,
	}
	return tempCell
}

func (c *Cells) GetMark() string {
	return string(c.mark)
}

func (c *Cells) SetMark(mark Mark) {
	c.mark = mark
}

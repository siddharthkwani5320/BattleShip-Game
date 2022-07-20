package component

type Player struct {
	nm     string
	noofHit  int
	noofMiss int
}

func NewPlayer(name string) *Player {
	var tempPlayer = &Player{
		nm:     name,
		noofHit:  0,
		noofMiss: 0,
	}
	return tempPlayer
}

func (p *Player) GetName() string {
	return p.nm
}

func (p *Player) GetHit() int {
	return p.noofHit
}

func (p *Player) GetMiss() int {
	return p.noofMiss
}

func (p *Player) IncrementHit() {
	p.noofHit++
}

func (p *Player) IncrementMiss() {
	p.noofMiss++
}

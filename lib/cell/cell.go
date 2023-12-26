package cell

import "github.com/mdwhatcott/go-set/v2/set"

type Cell struct {
	links set.Set[*Cell]

	Row int
	Col int

	North *Cell
	South *Cell
	East  *Cell
	West  *Cell
}

func New(row, col int) *Cell {
	return &Cell{links: make(set.Set[*Cell]), Row: row, Col: col}
}

func (this *Cell) IsLinkedTo(that *Cell) bool {
	return this.links.Contains(that)
}
func (this *Cell) LinkBoth(that *Cell) {
	this.links.Add(that)
	that.links.Add(this)
}
func (this *Cell) UnlinkBoth(that *Cell) {
	this.links.Remove(that)
	that.links.Remove(this)
}
func (this *Cell) Link(that *Cell) {
	this.links.Add(that)
}
func (this *Cell) Unlink(that *Cell) {
	this.links.Remove(that)
}

func (this *Cell) Neighbors() (result []*Cell) {
	if this.North != nil {
		result = append(result, this.North)
	}
	if this.South != nil {
		result = append(result, this.South)
	}
	if this.East != nil {
		result = append(result, this.East)
	}
	if this.West != nil {
		result = append(result, this.West)
	}
	return result
}

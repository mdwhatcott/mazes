package cell_test

import (
	"testing"

	"github.com/mdwhatcott/mazes/lib/cell"
	"github.com/mdwhatcott/testing/should"
)

func TestCellSuite(t *testing.T) {
	should.Run(&CellSuite{T: should.New(t)}, should.Options.UnitTests())
}

type CellSuite struct {
	*should.T
}

func (this *CellSuite) Setup() {
}

func (this *CellSuite) TestRowCol() {
	c := cell.New(1, 2)
	this.So(c.Row, should.Equal, 1)
	this.So(c.Col, should.Equal, 2)
}
func (this *CellSuite) TestLink() {
	c := cell.New(1, 2)
	d := cell.New(3, 4)

	this.So(c.IsLinkedTo(d), should.BeFalse)
	this.So(d.IsLinkedTo(c), should.BeFalse)
	c.Link(d)
	this.So(c.IsLinkedTo(d), should.BeTrue)
	this.So(d.IsLinkedTo(c), should.BeFalse)
	c.Unlink(d)
	this.So(c.IsLinkedTo(d), should.BeFalse)
	this.So(d.IsLinkedTo(c), should.BeFalse)
}
func (this *CellSuite) TestLinkBoth() {
	c := cell.New(1, 2)
	d := cell.New(3, 4)

	this.So(c.IsLinkedTo(d), should.BeFalse)
	this.So(d.IsLinkedTo(c), should.BeFalse)
	c.LinkBoth(d)
	this.So(c.IsLinkedTo(d), should.BeTrue)
	this.So(d.IsLinkedTo(c), should.BeTrue)
	c.UnlinkBoth(d)
	this.So(c.IsLinkedTo(d), should.BeFalse)
	this.So(d.IsLinkedTo(c), should.BeFalse)
}
func (this *CellSuite) TestNeighbors() {
	center := cell.New(3, 3)
	center.North = cell.New(2, 3)
	center.South = cell.New(4, 3)
	center.East = cell.New(3, 4)
	center.West = cell.New(4, 3)
	this.So(center.Neighbors(), should.Equal, []*cell.Cell{
		center.North, center.South, center.East, center.West,
	})

	center.North = nil
	this.So(center.Neighbors(), should.Equal, []*cell.Cell{
		center.South, center.East, center.West,
	})

	center.South = nil
	this.So(center.Neighbors(), should.Equal, []*cell.Cell{
		center.East, center.West,
	})

	center.East = nil
	this.So(center.Neighbors(), should.Equal, []*cell.Cell{
		center.West,
	})

	center.West = nil
	this.So(center.Neighbors(), should.BeEmpty)
}

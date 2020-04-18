package gopherplot

import (
	"image"
	"image/color"
)

//Renderable supports drawing
type Renderable interface {
	Draw(*image.RGBA)
}

//Projectable onto a 2-D surface
type Projectable interface {
	Project(ds *DataSpace) (X, Y int)
}

// Point defines a point
type Point struct {
	X float64
	Y float64
	Z float64
	Projectable
}

// Project the point given its position within the DataSpace
func (p *Point) Project(ds *DataSpace) (int, int) {
	var x, y int
	//TODO: figure this out
	return x, y
}

// Vertex is a point used in a line
type Vertex Point

/*
SimpleShape defines structures that can be drawns
*/
type SimpleShape interface {
	AppendVertex(Vertex)
	RemoveVertex(int)
	Projectable
}

/*
ClosedPolygon is a SimpleShape. The last vertex is assumed to connect to the first,
which "closes" the polygon
*/
type ClosedPolygon struct {
	Points   []Vertex
	IsFilled bool
	SimpleShape
}

// Line is the collection of points between two Verteces
type Line struct {
	P1    Vertex
	P2    Vertex
	Color color.RGBA
	Renderable
}

// Draw the line between points
func (l *Line) Draw(i *image.RGBA) {

}

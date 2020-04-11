package gopherplot

import "image/color"

// Label a string to identify a single item
type Label string

// Class a string to identify a group of items that share something in common
type Class string

// DataPoint is a datum
type DataPoint struct {
	X     float64
	Y     float64
	Z     float64
	Color color.RGBA
	Label
	Class
}

// Dimension X, Y, Z dimensional meta-data
type Dimension struct {
	Scale           float64
	MajorTicSpacing float64
	MinorTicSpacing float64
	ShowMinorTics   bool
	ShowMajorTics   bool
	TiltAngle       float64
}

// DataSpace is the renderable domain
type DataSpace struct {
	X, Y, Z    Dimension
	Data       []DataPoint
	Background color.RGBA
}

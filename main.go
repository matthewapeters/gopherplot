package gopherplot

import (
	"image"
	"image/color"
)

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
	X, Y, Z                   Dimension
	Data                      []DataPoint
	Background                color.RGBA
	RenderWidth, RenderHeight int
}

func (ds *DataSpace) checkDefaults() {
	if ds.RenderWidth == 0 {
		ds.RenderWidth = 800
	}
	if ds.RenderHeight == 0 {
		ds.RenderHeight = 640
	}
	allZeros := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	if ds.Background == allZeros {
		ds.Background = color.RGBA{255, 255, 255, 255}
	}

}

// Render returns a PNG image of the DataSpace
func (ds *DataSpace) Render() *image.Image {
	ds.checkDefaults()
	r := image.NewRGBA(image.Rect(0, 0, ds.RenderWidth, ds.RenderHeight))
	for y := 0; y < ds.RenderHeight; y++ {
		for x := 0; x < ds.RenderWidth; x++ {
			if x != y && y != ds.RenderWidth-x {
				r.Set(x, y, ds.Background)
			} else {

				r.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	i := image.Image(r)
	return &i
}

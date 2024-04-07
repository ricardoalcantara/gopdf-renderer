package pdfrenderer

type Color struct {
	R uint8
	G uint8
	B uint8
}

type Vec2 struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Rect struct {
	Position Vec2
	Size     Size
}

var RED = Color{R: 255, G: 0, B: 0}
var GREEN = Color{R: 0, G: 255, B: 0}
var BLUE = Color{R: 0, G: 0, B: 255}
var WHITE = Color{R: 255, G: 255, B: 255}
var BLACK = Color{R: 0, G: 0, B: 0}

type Alignment int

const (
	Start        Alignment = 0
	End          Alignment = 1
	Center       Alignment = 2
	SpaceBetween Alignment = 3
	SpaceEven    Alignment = 4
)

package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Line struct {
	elements []IElement
	area     Rect
	bgColor  *Color

	justifyContent Alignment
	alignItems     Alignment
}

func (p *Line) JustifyContent(alignment Alignment) *Line {
	p.justifyContent = alignment
	return p
}

func (p *Line) AlignItems(alignment Alignment) *Line {
	p.alignItems = alignment
	return p
}

func (p *Line) BgColor(color Color) *Line {
	p.bgColor = &color
	return p
}

func (p *Line) Size(width float64, height float64) *Line {
	p.area.Size.Width = width
	p.area.Size.Height = height
	return p
}

func (p *Line) Position(x float64, y float64) *Line {
	p.area.Position.X = x
	p.area.Position.Y = y
	return p
}

func (p *Line) GetArea() Rect {
	return p.area
}

func (p *Line) Text(value string) *Text {
	text := &Text{
		text: value,
	}

	p.elements = append(p.elements, text)
	return text
}

func (p *Line) Image(path string) *Image {
	img := &Image{
		path: path,
	}
	p.elements = append(p.elements, img)
	return img
}

func (l *Line) Draw(pdf *gopdf.GoPdf) {
	for _, element := range l.elements {
		size := element.Measure(pdf)
		if size.Height > l.area.Size.Height {
			l.area.Size.Height = size.Height
		}
	}

	if l.bgColor != nil {
		pdf.SetFillColor(l.bgColor.R, l.bgColor.G, l.bgColor.B)
		pdf.Rectangle(l.area.Position.X, l.area.Position.Y, l.area.Position.X+l.area.Size.Width, l.area.Position.Y+l.area.Size.Height, "F", 0, 0)
	}

	position := l.area.Position
	for _, element := range l.elements {
		pdf.SetXY(position.X, position.Y)
		size := element.Draw(pdf)
		if size.Height > l.area.Size.Height {
			l.area.Size.Height = size.Height
		}
		position.X += size.Width
	}
}

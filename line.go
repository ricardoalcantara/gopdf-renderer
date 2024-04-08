package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Line struct {
	elements []IElement
	area     Rect
	bgColor  *Color

	xAlignment Alignment
	yAlignment Alignment
}

func (p *Line) XAlignment(alignment Alignment) *Line {
	p.xAlignment = alignment
	return p
}

func (p *Line) YAlignment(alignment Alignment) *Line {
	p.yAlignment = alignment
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
	maxWidth := 0.0
	for _, element := range l.elements {
		size := element.Measure(pdf)
		maxWidth += size.Width
		if size.Height > l.area.Size.Height {
			l.area.Size.Height = size.Height
		}
	}

	if l.bgColor != nil {
		pdf.SetFillColor(l.bgColor.R, l.bgColor.G, l.bgColor.B)
		pdf.Rectangle(l.area.Position.X, l.area.Position.Y, l.area.Position.X+l.area.Size.Width, l.area.Position.Y+l.area.Size.Height, "F", 0, 0)
	}

	position := l.area.Position
	elementsCount := len(l.elements)
	switch l.xAlignment {
	case Start:
		for _, element := range l.elements {
			pdf.SetXY(position.X, position.Y)
			size := element.Draw(pdf)
			if size.Height > l.area.Size.Height {
				l.area.Size.Height = size.Height
			}
			position.X += size.Width
		}
	case SpaceEven:
		gap := (l.area.Size.Width - maxWidth) / float64(elementsCount)
		for _, element := range l.elements {
			size := element.Measure(pdf)
			pdf.SetXY(position.X+gap-size.Width/5, position.Y+(l.area.Size.Height/2)-(size.Height/2))
			element.Draw(pdf)
			position.X += size.Width
		}
	case SpaceBetween:
		for idx, element := range l.elements {
			size := element.Measure(pdf)
			if idx == 0 {
				pdf.SetXY(position.X, position.Y)
			} else if idx == elementsCount-1 {
				pdf.SetXY(l.area.Size.Width-size.Width, position.Y)
			}
			element.Draw(pdf)
			if size.Height > l.area.Size.Height {
				l.area.Size.Height = size.Height
			}
			position.X += size.Width
		}
	}
}

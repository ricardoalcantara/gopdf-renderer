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

	gap float64
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

func (p *Line) GetSize() Size {
	return p.area.Size
}

func (p *Line) GetPosition() Vec2 {
	return p.area.Position
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

func (p *Line) Table() *Table {
	table := &Table{
		size: Size{
			Width:  p.area.Size.Width,
			Height: p.area.Size.Height,
		},
	}
	p.elements = append(p.elements, table)
	return table
}

func (l *Line) UpdateMeasure(pdf *gopdf.GoPdf) {
	totalWidth := 0.0
	for _, element := range l.elements {
		element.Position(Vec2{
			X: l.area.Position.X + totalWidth,
			Y: l.area.Position.Y,
		})
		size := element.Measure(pdf)
		if size.Height > l.area.Size.Height {
			l.area.Size.Height = size.Height
		}
		totalWidth += size.Width
	}
}

func (l *Line) Update(pdf *gopdf.GoPdf) {
	l.UpdateMeasure(pdf)

	// Calculate x width
	totalElements := len(l.elements)
	position := l.area.Position
	switch l.xAlignment {
	case Start:
		for _, element := range l.elements {
			element.Position(Vec2{X: position.X, Y: position.Y})
			size := element.GetSize()
			position.X += size.Width + l.gap
		}
	case End:
		position.X = l.area.Position.X + l.area.Size.Width
		for i := totalElements - 1; i >= 0; i-- {
			element := l.elements[i]
			size := element.GetSize()
			position.X -= (size.Width + l.gap)
			element.Position(Vec2{X: position.X, Y: position.Y})
		}
	case Center:
		width := l.area.Size.Width / 2
		elementsSize := 0.0
		for _, element := range l.elements {
			elementsSize += element.GetSize().Width
		}
		position.X += width - (elementsSize / 2)
		for _, element := range l.elements {
			element.Position(Vec2{X: position.X, Y: position.Y})
			size := element.GetSize()
			position.X += size.Width + l.gap
		}
	case SpaceBetween:
		space := l.area.Size.Width / (float64(totalElements) + 1)
		for idx, element := range l.elements {
			size := element.GetSize()
			if idx == 0 {
				element.Position(Vec2{X: position.X, Y: position.Y})
			} else if idx == totalElements-1 {
				element.Position(Vec2{X: l.area.Size.Width - size.Width, Y: position.Y})
			} else {
				// TODO: I am not sure if this space / 2 is correct and it's not even tested
				position.X -= (size.Width / 2)
				element.Position(Vec2{X: position.X, Y: position.Y})
			}
			position.X += size.Width + (space / 2)
		}
	case SpaceEven:
		space := l.area.Size.Width / (float64(totalElements) + 1)
		position.X += space
		for _, element := range l.elements {
			size := element.GetSize()
			position.X -= (size.Width / 2)
			element.Position(Vec2{X: position.X, Y: position.Y})
			// TODO: I am not sure if this space / 2 is correct
			position.X += size.Width + (space / 2)
		}
	}

	// Calculate y height
	switch l.yAlignment {
	case Start:
		break
	case End:
		height := l.area.Size.Height
		for _, element := range l.elements {
			size := element.GetSize()
			element.Position(Vec2{X: element.GetPosition().X, Y: position.Y + height - size.Height})
		}
	case Center:
		centerHeight := l.area.Size.Height / 2
		for _, element := range l.elements {
			size := element.GetSize()
			element.Position(Vec2{X: element.GetPosition().X, Y: position.Y + centerHeight - (size.Height / 2)})
		}
	case SpaceBetween:
		break
	case SpaceEven:
		break
	}
}

func (l *Line) Draw(pdf *gopdf.GoPdf) {
	if l.bgColor != nil {
		pdf.SetFillColor(l.bgColor.R, l.bgColor.G, l.bgColor.B)
		pdf.Rectangle(
			l.area.Position.X,
			l.area.Position.Y,
			l.area.Position.X+l.area.Size.Width,
			l.area.Position.Y+l.area.Size.Height,
			"F", 0, 0)
	}

	for _, element := range l.elements {
		element.Draw(pdf)
	}
}

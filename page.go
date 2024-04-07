package pdfrenderer

import "github.com/signintech/gopdf"

type Page struct {
	lines        []*Line
	pageSize     gopdf.Rect
	marginTop    float64
	marginBottom float64
	marginLeft   float64
	marginRight  float64
}

func (p *Page) PageSize(pageSize gopdf.Rect) *Page {
	p.pageSize = pageSize
	return p
}

func (p *Page) MarginTop(marginTop float64) *Page {
	p.marginTop = marginTop
	return p
}

func (p *Page) MarginBottom(marginBottom float64) *Page {
	p.marginBottom = marginBottom
	return p
}

func (p *Page) MarginLeft(marginLeft float64) *Page {
	p.marginLeft = marginLeft
	return p
}

func (p *Page) MarginRight(marginRight float64) *Page {
	p.marginRight = marginRight
	return p
}

func (p *Page) MarginXY(x float64, y float64) *Page {
	p.marginTop = y
	p.marginBottom = y
	p.marginLeft = x
	p.marginRight = x
	return p
}

func (p *Page) Margin(value float64) *Page {
	p.marginTop = value
	p.marginBottom = value
	p.marginLeft = value
	p.marginRight = value
	return p
}

func (p *Page) Line(config func(line *Line)) *Line {
	line := &Line{
		area: Rect{
			Position: Vec2{X: p.marginLeft, Y: p.marginTop},
			Size:     Size{Width: p.pageSize.W - p.marginLeft - p.marginRight, Height: 0}},
	}

	if config != nil {
		config(line)
	}

	p.lines = append(p.lines, line)
	return line
}

func (p *Page) Draw(pdf *gopdf.GoPdf) float64 {
	lineHeight := p.marginTop
	for _, line := range p.lines {
		line.area.Position.Y = lineHeight
		line.Draw(pdf)
		lineHeight += line.area.Size.Height
	}

	return lineHeight
}

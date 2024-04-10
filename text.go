package pdfrenderer

import "github.com/signintech/gopdf"

type Text struct {
	text        string
	fontSize    float64
	color       Color
	size        Size
	position    Vec2
	borderColor *Color
}

func (t *Text) Position(position Vec2) {
	t.position = position
}

func (t *Text) GetPosition() Vec2 {
	return t.position
}

func (t *Text) FontSize(fontSize float64) *Text {
	t.fontSize = fontSize
	return t
}

func (t *Text) GetSize() Size {
	return t.size
}

func (t *Text) Color(color Color) *Text {
	t.color = color
	return t
}

func (t *Text) Measure(pdf *gopdf.GoPdf) Size {
	pdf.SetFontSize(t.fontSize)
	width, err := pdf.MeasureTextWidth(t.text)
	if err != nil {
		panic(err)
	}
	height, err := pdf.MeasureCellHeightByText(t.text)
	if err != nil {
		panic(err)
	}

	t.size = Size{
		Width:  width,
		Height: height,
	}

	return t.size
}

func (t Text) Draw(pdf *gopdf.GoPdf) {
	if t.borderColor != nil {
		pdf.SetFillColor(t.borderColor.R, t.borderColor.G, t.borderColor.B)
		pdf.SetLineWidth(1.0)
		pdf.Rectangle(
			t.position.X,
			t.position.Y,
			t.position.X+t.size.Width,
			t.position.Y+t.size.Height,
			"D", 0, 0)
	}

	pdf.SetFillColor(t.color.R, t.color.G, t.color.B)
	pdf.SetFontSize(t.fontSize)
	pdf.SetXY(t.position.X, t.position.Y+t.size.Height-2.5)

	err := pdf.Text(t.text)
	if err != nil {
		panic(err)
	}
}

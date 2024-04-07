package pdfrenderer

import "github.com/signintech/gopdf"

type Text struct {
	text  string
	size  float64
	color Color
	// Area  Rect
}

func (t *Text) Size(size float64) *Text {
	t.size = size
	return t
}

func (t *Text) Color(color Color) *Text {
	t.color = color
	return t
}

func (t *Text) Measure(pdf *gopdf.GoPdf) Size {
	pdf.SetFontSize(t.size)
	width, err := pdf.MeasureTextWidth(t.text)
	if err != nil {
		panic(err)
	}
	height, err := pdf.MeasureCellHeightByText(t.text)
	if err != nil {
		panic(err)
	}

	return Size{
		Width:  width,
		Height: height,
	}
}

func (t Text) Draw(pdf *gopdf.GoPdf) Size {
	pdf.SetFillColor(t.color.R, t.color.G, t.color.B)
	pdf.SetFontSize(t.size)

	size := t.Measure(pdf)
	pdf.SetY(pdf.GetY() + size.Height - 2.5)

	err := pdf.Text(t.text)
	if err != nil {
		panic(err)
	}

	return size
}

package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Table struct {
	size     Size
	position Vec2
	header   []string
	rows     [][]string
}

func (t *Table) Position(position Vec2) {
	t.position = position
}

func (t *Table) GetPosition() Vec2 {
	return t.position
}

func (t *Table) GetSize() Size {
	return t.size
}

func (t *Table) Header(header []string) *Table {
	t.header = header
	return t
}

func (t *Table) Rows(rows [][]string) *Table {
	t.rows = rows
	return t
}

func (t *Table) AddRow(rows []string) *Table {
	t.rows = append(t.rows, rows)
	return t
}

func (t *Table) Measure(pdf *gopdf.GoPdf) Size {
	return t.size
}

func (t *Table) Draw(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(BLUE.R, BLUE.G, BLUE.B)
	pdf.SetFontSize(12)
	pdf.SetXY(t.position.X, t.position.Y+t.size.Height-2.5)

	err := pdf.Text("Hello World")
	if err != nil {
		panic(err)
	}
}

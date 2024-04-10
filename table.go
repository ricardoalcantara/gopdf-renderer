package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Table struct {
	size     Size
	position Vec2
	rows     [][]Text
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

func (t *Table) Rows(rows [][]string) *Table {
	for _, r := range rows {
		t.AddRow(r)
	}
	return t
}

func (t *Table) AddRow(row []string) *Table {
	r := make([]Text, len(row))
	for idx, h := range row {
		r[idx] = Text{
			text:        h,
			fontSize:    12,
			color:       BLACK,
			borderColor: &BLACK,
		}
	}

	t.rows = append(t.rows, r)
	return t
}

func (t *Table) Measure(pdf *gopdf.GoPdf) Size {
	rowsCount := len(t.rows)
	if rowsCount == 0 {
		return Size{}
	}
	header := t.rows[0]

	cellWidth := t.size.Width / float64(len(header))

	for _, row := range t.rows {
		cellHeight := 0.0
		for idx := range row {
			cell := &row[idx]
			size := cell.Measure(pdf)
			if size.Height > cellHeight {
				cellHeight = size.Height 
			}
			cell.size.Width = cellWidth
		}
		cellHeight += 5
		t.size.Height += cellHeight
		for idx := range row {
			cell := &row[idx]
			cell.size.Height = cellHeight
			cell.position.X = t.position.X + (cellWidth * float64(idx))
			cell.position.Y = t.position.Y + t.size.Height
		}

	}
	return t.size
}

func (t *Table) Draw(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(BLUE.R, BLUE.G, BLUE.B)
	pdf.SetFontSize(12)

	for _, row := range t.rows {
		for _, cell := range row {
			cell.Draw(pdf)
		}
	}

	// cellWidth := t.size.Width / float64(len(header))

	// for idx, header := range header {
	// 	// width, err := pdf.MeasureTextWidth(header)
	// 	height, err := pdf.MeasureCellHeightByText(header)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	x := t.position.X + (cellWidth * float64(idx))
	// 	y := t.position.Y + height - 2.5
	// 	pdf.SetXY(x+2.5, y)
	// 	pdf.SetFillColor(BLACK.R, BLACK.G, BLACK.B)
	// 	pdf.Rectangle(
	// 		x,
	// 		y-height,
	// 		x+cellWidth,
	// 		y+2.5,
	// 		"D", 0, 0)
	// 	err = pdf.Text(header)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// for _, row := range t.rows {
	// 	for idx, cell := range row {
	// 		// width, err := pdf.MeasureTextWidth(header)
	// 		height, err := pdf.MeasureCellHeightByText(cell)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		x := t.position.X + (cellWidth * float64(idx))
	// 		y := t.position.Y + height - 2.5
	// 		pdf.SetXY(x+2.5, y)
	// 		pdf.SetFillColor(BLACK.R, BLACK.G, BLACK.B)
	// 		pdf.Rectangle(
	// 			x,
	// 			y-height,
	// 			x+cellWidth,
	// 			y+2.5,
	// 			"D", 0, 0)
	// 		err = pdf.Text(cell)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
	// pdf.SetXY(t.position.X, t.position.Y+t.size.Height-2.5)
	// err := pdf.Text("Hello World")
	// if err != nil {
	// 	panic(err)
	// }
}

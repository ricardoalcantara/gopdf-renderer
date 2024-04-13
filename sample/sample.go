package main

import (
	"io/ioutil"

	pdfrenderer "github.com/ricardoalcantara/gopdf-renderer"
	"github.com/rs/zerolog/log"
)

var primary = pdfrenderer.Color{R: 200, G: 200, B: 200}

func main() {
	pdf := pdfrenderer.NewPdfRenderer().
		PageSize(pdfrenderer.PageSizeA4).
		LoadFont("DejaVuSans", "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf")

	pdf.
		Header(func(page *pdfrenderer.Page) {
			page.
				Margin(30)

			page.
				Line(func(line *pdfrenderer.Line) {
					line.
						XAlignment(pdfrenderer.SpaceBetween)
					line.Image("./sample/github.png").
						Size(32, 32)
					line.
						Text("Sample Company").
						FontSize(10)
				})
		})

	pdf.
		Footer(func(page *pdfrenderer.Page) {
			page.
				Margin(30)

			page.
				Line(func(line *pdfrenderer.Line) {
					line.
						XAlignment(pdfrenderer.Center)
					line.
						Text("© 2024 Sample Company. All Rights Reserved").
						FontSize(10)
				})
			page.
				Line(func(line *pdfrenderer.Line) {
					line.
						XAlignment(pdfrenderer.Center)
					line.
						Text("123 Main Street, Sampletown, USA").
						FontSize(10)
				})
		})

	page := pdf.
		Page().
		Margin(30)

	page.
		Line(func(line *pdfrenderer.Line) {
			line.
				BgColor(primary).
				Size(500, 50).
				Position(0, line.GetPosition().Y).
				XAlignment(pdfrenderer.SpaceEven).
				YAlignment(pdfrenderer.Center)

			line.
				Text("Invoice #123456789").
				Color(pdfrenderer.WHITE).
				FontSize(21)
			line.
				Text("Paid").
				Color(pdfrenderer.WHITE).
				FontSize(21)
		})

	page.BreakLine(10)
	page.
		Line(func(line *pdfrenderer.Line) {
			line.Text("Billed to").
				FontSize(10)
		})
	page.BreakLine(10)
	page.
		Line(func(line *pdfrenderer.Line) {
			line.Text("John Doe").
				FontSize(10)
		})
	page.
		Line(func(line *pdfrenderer.Line) {
			line.Text("123 Maple Street").
				FontSize(10)
		})
	page.
		Line(func(line *pdfrenderer.Line) {
			line.Text("Anytown, USA, 12345").
				FontSize(10)
		})
	page.BreakLine(10)
	page.
		Line(func(line *pdfrenderer.Line) {
			line.Text("Billing Period: April 2024").
				FontSize(10)
		})
	page.BreakLine(20)

	page.
		Line(nil).
		Table().
		AddRow([]string{"Description", "Quantity", "Total"}).
		Rows([][]string{
			{"Dedicated Server", "1", "$1400.00"},
			{"Additional IP", "13", "$208.00"},
		}).
		AddRow([]string{"", "Total", "$1608.00"})

	err := pdf.Render("invoice.pdf")

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	b, err := pdf.RenderBytes()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	err = ioutil.WriteFile("invoice2.pdf", b, 0644)
	if err != nil {
		panic(err)
	}
}

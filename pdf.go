package pdfrenderer

import "github.com/signintech/gopdf"

type PdfRenderer struct {
	header   *Page
	footer   *Page
	pages    []*Page
	pageSize Size
}

func NewPdfRenderer() *PdfRenderer {
	return &PdfRenderer{}
}

func (p *PdfRenderer) LoadFont(name string, path string) *PdfRenderer {
	return p
}

func (p *PdfRenderer) PageSize(pageSize Size) *PdfRenderer {
	p.pageSize = pageSize
	return p
}

func (p *PdfRenderer) Header(config func(page *Page)) *Page {
	p.header = &Page{
		pageSize: p.pageSize,
	}
	if config != nil {
		config(p.header)
	}
	return p.header
}

func (p *PdfRenderer) Footer(config func(page *Page)) *Page {
	p.footer = &Page{
		pageSize: p.pageSize,
	}
	if config != nil {
		config(p.footer)
	}
	return p.footer
}

func (p *PdfRenderer) Page() *Page {
	page := &Page{
		pageSize: p.pageSize,
	}
	p.pages = append(p.pages, page)
	return page
}

func (p *PdfRenderer) Render(path string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: p.pageSize.Width, H: p.pageSize.Height}})

	var height float64
	if p.header != nil {
		pdf.AddHeader(func() {
			height = p.header.Draw(&pdf)
		})
	}

	if p.footer != nil {
		pdf.AddFooter(func() {
			p.footer.Draw(&pdf)
		})
	}

	err := pdf.AddTTFFont("DejaVuSans", "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf")

	if err != nil {
		return err
	}
	err = pdf.SetFont("DejaVuSans", "", 12)
	if err != nil {
		return err
	}

	for _, page := range p.pages {
		pdf.AddPage()
		pp := *page
		pp.marginTop += height
		pp.Draw(&pdf)
	}

	return pdf.WritePdf(path)
}

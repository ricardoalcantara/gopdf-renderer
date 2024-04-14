package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type PdfRenderer struct {
	header      *Page
	footer      *Page
	pages       []*Page
	pageSize    Size
	fonts       map[string]string
	defaultFont string
}

func NewPdfRenderer() *PdfRenderer {
	return &PdfRenderer{
		fonts: make(map[string]string),
	}
}

func (p *PdfRenderer) LoadFont(name string, path string) *PdfRenderer {
	if p.defaultFont == "" {
		p.defaultFont = name
	}
	p.fonts[name] = path
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
	pdf, err := p.build()
	if err != nil {
		return err
	}
	return pdf.WritePdf(path)
}

func (p *PdfRenderer) RenderBytes() ([]byte, error) {
	pdf, err := p.build()
	if err != nil {
		return nil, err
	}
	stream := MemoryStream{}
	pdf.WriteTo(&stream)
	return stream.Bytes(), nil
}

func (p *PdfRenderer) build() (*gopdf.GoPdf, error) {
	// BUG: if you build twice the values will be overwritten
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
			pdf.SetY(825)
			p.footer.DrawReverse(&pdf)
		})
	}

	for k, v := range p.fonts {
		err := pdf.AddTTFFont(k, v)
		if err != nil {
			return nil, err
		}
	}
	err := pdf.SetFont(p.defaultFont, "", 0)
	if err != nil {
		return nil, err
	}

	for _, page := range p.pages {
		pdf.AddPage()
		pp := *page
		pp.marginTop += height
		pp.Draw(&pdf)
	}

	return &pdf, nil
}

package pdfrenderer

import "github.com/signintech/gopdf"

type IElement interface {
	Measure(pdf *gopdf.GoPdf) Size
	Draw(*gopdf.GoPdf) Size
}

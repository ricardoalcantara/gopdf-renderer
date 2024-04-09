package pdfrenderer

import "github.com/signintech/gopdf"

type IElement interface {
	GetSize() Size
	GetPosition() Vec2
	Measure(pdf *gopdf.GoPdf) Size
	Position(position Vec2)
	Draw(*gopdf.GoPdf)
}

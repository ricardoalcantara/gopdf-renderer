package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Image struct {
	size     Size
	path     string
	position Vec2
}

func (i *Image) Position(position Vec2) {
	i.position = position
}

func (i *Image) GetPosition() Vec2 {
	return i.position
}

func (i *Image) Size(width float64, height float64) *Image {
	i.size = Size{
		Width:  width,
		Height: height,
	}
	return i
}

func (i *Image) GetSize() Size {
	return i.size
}

func (i *Image) Measure(pdf *gopdf.GoPdf) Size {

	imgh, err := gopdf.ImageHolderByPath(i.path)
	if err != nil {
		panic(err)
	}

	imgobj := new(gopdf.ImageObj)
	err = imgobj.SetImage(imgh)
	if err != nil {
		panic(err)
	}

	if i.size.Width == 0 || i.size.Height == 0 {
		rect := imgobj.GetRect()

		i.size.Width = rect.H
		i.size.Height = rect.W
	}

	return i.size
}

func (i *Image) Draw(pdf *gopdf.GoPdf) {
	pdfSize := &gopdf.Rect{W: i.size.Width, H: i.size.Height}
	err := pdf.Image(i.path, i.position.X, i.position.Y, pdfSize)
	if err != nil {
		panic(err)
	}
}

package pdfrenderer

import (
	"github.com/signintech/gopdf"
)

type Image struct {
	size Size
	path string
}

func (i *Image) Size(width float64, height float64) *Image {
	i.size = Size{
		Width:  width,
		Height: height,
	}
	return i
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

	var size Size
	if i.size.Width == 0 || i.size.Height == 0 {
		rect := imgobj.GetRect()

		size.Width = rect.H
		size.Height = rect.W
	}

	return size
}

func (i *Image) Draw(pdf *gopdf.GoPdf) Size {
	var pdfSize *gopdf.Rect
	var size Size
	if i.size.Width > 0 && i.size.Height > 0 {
		size = i.size
	} else {
		size = i.Measure(pdf)
	}
	pdfSize = &gopdf.Rect{W: size.Width, H: size.Height}

	err := pdf.Image(i.path, pdf.GetX(), pdf.GetY(), pdfSize)
	if err != nil {
		panic(err)
	}

	return size
}

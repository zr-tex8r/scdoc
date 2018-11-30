package scnix

import (
	"image/color"
	"io"
)

// Doc is a no-op scdoc.Doc.
type Doc struct {
	width, height float64
}

func (d *Doc) SetPageSize(width, height float64) error {
	d.width, d.height = width, height
	return nil
}

func (d *Doc) PageSize() (float64, float64) {
	return d.width, d.height
}

func (d *Doc) AddPageScaled(muffler color.Color, scale float64) error {
	return nil
}

func (d *Doc) AddPage(muffler color.Color) error {
	return nil
}

func (d *Doc) SetDocInfo(info map[string]string) (err error) {
	return nil
}

func (d *Doc) WriteTo(wr io.Writer) (n int64, err error) {
	return 0, nil
}

func (d *Doc) Bytes() (bs []byte, err error) {
	return nil, nil
}

// Copyright (c) 2018 Takayuki YATO (aka. "ZR")
//   GitHub:   https://github.com/zr-tex8r
//   Twitter:  @zr_tex8r
// Distributed under the MIT License.

// Package scdoc defines interface Doc, which represents SC-oriented
// documents of any format.
package scdoc

import (
	"image/color"
	"io"
)

// Doc represents an SC-oriented document of any format.
type Doc interface {
	// SetPageSize sets the width and the height (measured in PDF points)
	// of this document.
	SetPageSize(width, height float64) error
	// PageSize returns the width and height of this document.
	PageSize() (float64, float64)
	// AddPageScaled adds to this document a new page with the given
	// muffler color and the scale value to the standard snowman size.
	AddPageScaled(muffler color.Color, scale float64) error
	// AddPageScaled adds to this document a new page with the given
	// muffler color (without scale).
	AddPage(muffler color.Color) error
	// SetDocInfo specifies some information of this document.
	// The input is given as a map of strings.
	SetDocInfo(info map[string]string) (err error)
	// WriteTo generates the file content of this document and writes
	// the resulted bytes to an io.Writer.
	WriteTo(wr io.Writer) (n int64, err error)
	// Bytes generates the file content and returns the resulted bytes.
	// The same notice on WriteTo also applies.
	Bytes() (bs []byte, err error)
}

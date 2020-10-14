package puzzle

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type encoder interface {
	Encode(w io.Writer, m image.Image, o *jpeg.Options) error
}

type jpgEncoder struct {
}

func newJpgEncoder() *jpgEncoder {
	return &jpgEncoder{}
}

func (j jpgEncoder) Encode(w io.Writer, m image.Image, o *jpeg.Options) error {
	return jpeg.Encode(w, m, o)
}

type pngEncoder struct {
}

func newPngEncoder() *pngEncoder {
	return &pngEncoder{}
}

func (j pngEncoder) Encode(w io.Writer, m image.Image, o *jpeg.Options) error {
	enc := png.Encoder{
		CompressionLevel: png.BestSpeed,
		BufferPool:       nil,
	}
	return enc.Encode(w, m)
}

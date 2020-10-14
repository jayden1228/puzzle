package puzzle

import (
	"image"
	"log"
	"os"
	"path"
)

func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func SaveImage(p string, m image.Image) error {
	var encoder encoder
	ext := path.Ext(p)
	switch ext {
	case "jpg", "jpeg":
		encoder = newJpgEncoder()
	case "png":
		encoder = newPngEncoder()
	default:
		encoder = newPngEncoder()
	}

	out, err := os.Create(p)

	if err != nil {
		log.Printf("Error save image file: %+v\n", err)
		return err
	}

	return encoder.Encode(out, m, nil)
}

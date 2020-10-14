package main

import (
	"image/color"
	_ "image/png"
	"log"

	m "github.com/jayden1228/puzzle"
)

func main() {

	assetPath := "./assets/"
	frameFileNames := []string{"frame_1.png", "frame_2.png", "frame_3.png"}

	frameOffset := 0
	imgs := make([]m.ImageLayer, 0)
	for i, _ := range frameFileNames {
		framePath := assetPath + frameFileNames[i]
		img, err := m.LoadImage(framePath)
		if err != nil {
			log.Println(err)
			return
		}

		imgs = append(imgs, m.ImageLayer{
			Image: img,
			XPos:  frameOffset,
			YPos:  0,
		})

		frameOffset = frameOffset + img.Bounds().Size().X
	}

	bg := m.BgProperty{
		Width:   frameOffset,
		Height:  imgs[0].Image.Bounds().Size().Y,
		BgColor: color.RGBA{},
	}

	res, err := m.GenerateImage(imgs, nil, bg)
	if err != nil {
		log.Printf("Error generating image: %+v\n", err)
		return
	}

	outputPath := "./output.png"
	err = m.SaveImage(outputPath, res)
	if err != nil {
		log.Printf("Error generating image: %+v\n", err)
		return
	}
}

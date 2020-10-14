package main

import (
	"image/color"
	_ "image/png"
	"log"

	m "github.com/jayden1228/puzzle"
)

func main() {
	inputPath := "./assets/images/certificate.jpg"
	outputPath := "./output.png"
	bgImage, err := m.LoadImage(inputPath)
	if err != nil {
		log.Println(err)
		return
	}

	imgs := []m.ImageLayer{
		m.ImageLayer{
			Image: bgImage,
			XPos:  0,
			YPos:  0,
		},
	}

	bg := m.BgProperty{
		Width:   bgImage.Bounds().Size().X,
		Height:  bgImage.Bounds().Size().Y,
		BgColor: color.RGBA{},
	}

	DefaultColor := color.Black
	DeafaultFontPath := "./assets/fonts/"
	DefaultFont := "msyh.ttf"

	texts := []m.TextLayer{
		m.TextLayer{
			Content:  "张小鹏",
			FontPath: DeafaultFontPath,
			FontType: DefaultFont,
			Size:     11,
			Color:    DefaultColor,
			XPos:     316,
			YPos:     424,
		},
		m.TextLayer{
			Content:  "鉴于您在“造物编程课” 入门课 课程中完成全部项目",
			FontPath: DeafaultFontPath,
			FontType: DefaultFont,
			Size:     5,
			Color:    DefaultColor,
			XPos:     194,
			YPos:     500,
		},
		m.TextLayer{
			Content:  "2020.08.20",
			FontPath: DeafaultFontPath,
			FontType: DefaultFont,
			Size:     4,
			Color:    DefaultColor,
			XPos:     572,
			YPos:     1012,
		},
	}

	res, err := m.GenerateImage(imgs, texts, bg)
	if err != nil {
		log.Printf("Error generating image: %+v\n", err)
		return
	}

	err = m.SaveImage(outputPath, res)
	if err != nil {
		log.Printf("Error generating image: %+v\n", err)
		return
	}
}

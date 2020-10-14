package puzzle

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)

type (
	ImageLayer struct {
		Image image.Image
		XPos  int
		YPos  int
	}

	TextLayer struct {
		Content  string
		FontPath string
		FontType string
		Size     float64
		Color    color.Color
		DPI      float64
		Spacing  float64
		XPos     int
		YPos     int
	}

	BgProperty struct {
		Width   int
		Height  int
		BgColor color.Color
	}
)

// GenerateImage combine images and texts into one image
func GenerateImage(images []ImageLayer, texts []TextLayer, bgProperty BgProperty) (*image.RGBA, error) {
	// create image's background
	bgImg := image.NewRGBA(image.Rect(0, 0, bgProperty.Width, bgProperty.Height))

	// set the background color
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{C: bgProperty.BgColor}, image.Point{}, draw.Src)

	// looping image layer, higher array index = upper layer
	for _, img := range images {
		//set image offset
		offset := image.Pt(img.XPos, img.YPos)

		//combine the image
		draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.Image, image.Point{}, draw.Over)
	}

	// add text(s)
	bgImg, err := addText(bgImg, texts)
	if err != nil {
		return nil, err
	}

	return bgImg, nil
}

var (
	defaultColor         = color.Black
	defaultDPI   float64 = 256
	defaultSpace float64 = 1.5
)

func addText(img *image.RGBA, texts []TextLayer) (*image.RGBA, error) {
	//initialize the context
	c := freetype.NewContext()

	for _, text := range texts {
		// fill default property
		if text.Color == nil {
			text.Color = defaultColor
		}
		if text.DPI == 0 {
			text.DPI = defaultDPI
		}
		if text.Spacing == 0 {
			text.Spacing = defaultSpace
		}
		//read font data
		//fullPath := path.Join(text.FontPath, text.FontType)
		fontBytes, err := ioutil.ReadFile(text.FontPath + text.FontType)
		if err != nil {
			return nil, err
		}
		f, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return nil, err
		}

		//set text configuration
		c.SetDPI(text.DPI)
		c.SetFont(f)
		c.SetFontSize(text.Size)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(image.NewUniform(text.Color))

		//positioning the text
		pt := freetype.Pt(text.XPos, text.YPos+int(c.PointToFixed(text.Size)>>6))

		//draw the text on image
		_, err = c.DrawString(text.Content, pt)
		if err != nil {
			log.Println(err)
			return img, nil
		}
		pt.Y += c.PointToFixed(text.Size * text.Spacing)
	}

	return img, nil
}

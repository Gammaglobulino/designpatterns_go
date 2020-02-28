package shapes

import (
	"../strategy"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
)

type ImageSquare struct {
	strategy.PrintOutput
}

func (im *ImageSquare) Print() error {
	width := 800
	height := 600
	origin := image.Point{
		X: 0,
		Y: 0,
	}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{width, height},
	})
	var bgColor = image.Uniform{color.RGBA{
		R: 70,
		G: 70,
		B: 70,
		A: 0,
	}}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 1,
	}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)
	if im.Writer == nil {
		return fmt.Errorf("No writer stored on ImageSquare")
	}
	if err := jpeg.Encode(im.Writer, bgImage, quality); err != nil {
		return fmt.Errorf("Error saving the image to the file")
	}
	if im.LogWriter != nil {
		io.Copy(im.LogWriter, bytes.NewReader([]byte("Image written in provided writer\n")))
	}
	return nil
}

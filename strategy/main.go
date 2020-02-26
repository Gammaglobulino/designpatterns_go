package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type PrintStrategy interface {
	Print() error
}

type TextSquare struct {
}

type ImageSquare struct {
	DestinationFilePath string
}

func (tx *TextSquare) Print() error {
	fmt.Println("Square")
	return nil
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

	w, err := os.Create(im.DestinationFilePath)
	if err != nil {
		return err
	}
	defer w.Close()
	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("Error saving the image to the file")
	}
	return nil
}

func main() {
	var output = flag.String("output", "console", "The output to use between 'console' and 'image' file ")
	flag.Parse()
	var activeStrategy PrintStrategy
	switch *output {
	case "console":
		activeStrategy = &TextSquare{}
	case "image":
		activeStrategy = &ImageSquare{DestinationFilePath: "image.jpg"}
	default:
		activeStrategy = &TextSquare{}
	}
	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}

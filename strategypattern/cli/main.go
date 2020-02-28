package main

import (
	"../shapes"
	"flag"
	"log"
	"os"
)

func main() {
	var output = flag.String("output", "text", "The output to use between 'console' and 'image' file ")
	flag.Parse()
	activeStrategy, err := shapes.NewPrinter(*output)
	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("image.jpg")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	default:
		activeStrategy = &shapes.TextSquare{}
	}
	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}

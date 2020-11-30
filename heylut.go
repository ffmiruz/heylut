package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ./heylut <image>")
	}

	subject, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer subject.Close()

	image, err := jpeg.Decode(subject)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("lut")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		lut, err := ioutil.ReadFile("lut/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		result := applyLutToImage(image, lut)

		out, err := os.Create(file.Name() + "_test.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		err = jpeg.Encode(out, result, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func applyLutToImage(input image.Image, lut []byte) image.Image {
	result := imaging.AdjustFunc(input, func(c color.NRGBA) color.NRGBA {
		c.R = uint8(lut[c.R])
		c.G = uint8(lut[c.G])
		c.B = uint8(lut[c.B])
		return c
	})
	return result
}

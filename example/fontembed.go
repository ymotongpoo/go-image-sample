package main

import (
	"fmt"
	"image"
	_ "image/draw"
	_ "image/gif"
	"image/jpeg"
	"image/color"
	"io/ioutil"
	_ "log"
	"os"

	"code.google.com/p/freetype-go/truetype"
)

func main() {
	// Open source animation GIF file.
	file, err := os.Open("image/burger.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	j, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Prepare font data.
	FontBytes, err := ioutil.ReadFile("font/mikachan.ttf")
	if err != nil {
		panic(err)
	}
	font, err := truetype.ParseFont(FontBytes)
	if err != nil {
		panic(err)
	}

	image.Uniform(color.

	text := "バーガー食べやがれ！"

	// Dump image data into file.
	file, err = os.Create("image/new-burger.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = jpeg.Encode(file, dst)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote out new-burger.jpeg")
}

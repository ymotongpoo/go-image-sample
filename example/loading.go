package main

import (
	"fmt"
	"image"
	_ "image/draw"
	"image/gif"
	_ "log"
	"os"
)

func main() {
	// Open source animation GIF file.
	file, err := os.Open("image/loading.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}

	// Initialize destination animation GIF
	dst := &gif.GIF{
		Image:     make([]*image.Paletted, len(g.Image)/2),
		Delay:     make([]int, len(g.Image)/2),
		LoopCount: 0,
	}

	for i, img := range g.Image {
		if i%2 == 0 {
			dst.Image[i/2] = img
		}
	}

	// Dump image data into file.
	file, err = os.Create("image/fast-loading.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gif.EncodeAll(file, dst)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote out fast-loading.gif")
}

package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"os"
)

func main() {
	// Open source PNG file.
	file, err := os.Open("image/gopher.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// Define destination boundary. Expecting original image is square.
	r := data.Bounds()

	// Prepare distination image buffer.
	dst := gif.GIF{
		Image: []*image.Paletted{},
	}

	// Rotate original image and store them into destination.
	original := image.NewPaletted(r, palette.WebSafe)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			original.Set(x, y, data.At(x, y))
		}
	}
	dst.Image = append(dst.Image, original)

	clockwise := image.NewPaletted(r, palette.WebSafe)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			clockwise.Set(x, y, data.At(-y+r.Max.Y, x))
		}
	}
	dst.Image = append(dst.Image, clockwise)

	upsidedown := image.NewPaletted(r, palette.WebSafe)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			upsidedown.Set(x, y, data.At(-y+r.Max.Y, -x+r.Max.X))
		}
	}
	dst.Image = append(dst.Image, upsidedown)

	counterclockwise := image.NewPaletted(r, palette.WebSafe)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			counterclockwise.Set(x, y, data.At(x, -y+r.Max.X))
		}
	}
	dst.Image = append(dst.Image, counterclockwise)

	// Post process
	dst.Delay = make([]int, len(dst.Image))
	dst.LoopCount = 100

	// Dump image data into file.
	file, err = os.Create("image/rotate-gopher.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gif.EncodeAll(file, &dst)
	if err != nil {
		panic(err)
	}
	log.Println("wrote out rotate-gopher.gif")
}

package main

/*
** 15.05.2019, Sergei Khamenev, Task-1: Tests
 */

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var white color.Color = color.RGBA{251, 252, 254, 255}

var black color.Color = color.RGBA{0, 1, 0, 255}

func main() {
	file, err := os.Create("image.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	// Белый фон
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	// Горизонтальная черная линия
	for x := 0; x < 500; x++ {
		y := 250
		img.Set(x, y, black)
	}

	// Вертикальная черная линия
	for y := 0; y < 500; y++ {
		x := 250
		img.Set(x, y, black)
	}

	png.Encode(file, img)
}

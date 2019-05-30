package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const dimm = 16

type hist [dimm][dimm][dimm]uint

type histrec struct {
	h *hist
	n string
}

func histogram(name string) (histrec, error) {
	fd, err := os.Open(name)
	if err != nil {
		panic(fmt.Sprintf("Cannot open file: %v\n", name))
	}
	defer fd.Close()

	var img image.Image
	switch filepath.Ext(name) {
	case ".jpg", ".jpeg", ".JPG":
		img, err = jpeg.Decode(fd)
		if err != nil {
			panic(fmt.Sprintf("Error at decoding: %v\n", name))
		}
	case ".png":
		img, err = png.Decode(fd)
		if err != nil {
			panic(fmt.Sprintf("Error at decoding: %v\n", name))
		}
	default:
		return histrec{}, fmt.Errorf("Cannot process %v", name)
	}

	var histogram hist

	boundaries := img.Bounds()
	for i := boundaries.Min.X; i < boundaries.Max.X; i++ {
		for j := boundaries.Min.Y; j < boundaries.Max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			histogram[r/(65536/dimm)][g/(65536/dimm)][b/(65536/dimm)]++
		}
	}

	return histrec{h: &histogram, n: name}, nil
}

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
)

const dimm = 16

type hist [dimm][dimm][dimm]uint

type histrec struct {
	h *hist
	n string
}

func histogram(name string) histrec {
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
		panic(fmt.Sprintf("Cannot decode: %v\n", name))
	}

	var histogram hist

	boundaries := img.Bounds()
	for i := boundaries.Min.X; i < boundaries.Max.X; i++ {
		for j := boundaries.Min.Y; j < boundaries.Max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			histogram[r/(65536/dimm)][g/(65536/dimm)][b/(65536/dimm)]++
		}
	}

	return histrec{h: &histogram, n: name}
}

func manhatanDistance(h1, h2 *hist) uint {
	var sum uint
	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				sum += uint(math.Abs(float64(h1[i][j][k]) - float64(h2[i][j][k])))
			}
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		panic("provide dir")
	}
	dir, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	hList := []histrec{}
	for _, f := range files {
		fmt.Printf("Processing: %v\n", os.Args[1]+"/"+f.Name())
		hList = append(hList, histogram(os.Args[1]+"/"+f.Name()))
	}

	for i := 0; i < len(hList)-1; i++ {
		for j := i + 1; j < len(hList); j++ {
			distance := manhatanDistance(hList[i].h, hList[j].h)
			if distance < 100000 {
				fmt.Printf("Similar: %v <-> %v  dist: %v\n", hList[i].n, hList[j].n, distance)
			}
		}
	}
}

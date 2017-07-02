package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var distanceFunction = flag.String("d", "manhattan", "distance type, default manhattan")
var toleration = flag.Uint("t", 100000, "distance under this value indicates that images are similar, default: 100000")
var path = flag.String("p", "", "path to folder with images")
var verbose = flag.Bool("v", false, "turn on verbose messages")

func main() {
	flag.Parse()

	if *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir, err := filepath.Abs(*path)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	hList := []histrec{}
	numFiles := len(files)
	for i, f := range files {
		if *verbose {
			fmt.Printf("Processing(%v/%v): %v\n", i, numFiles, *path+"/"+f.Name())
		} else {
			fmt.Printf("\r%d/%d", i, numFiles)
		}
		hList = append(hList, histogram(*path+"/"+f.Name()))
	}

	fmt.Println()
	distFunc := getDistanceFunction(*distanceFunction)
	for i := 0; i < len(hList)-1; i++ {
		for j := i + 1; j < len(hList); j++ {
			distance := distFunc(hList[i].h, hList[j].h)
			if distance < *toleration {
				fmt.Printf("Similar: %v <-> %v  dist: %v\n", hList[i].n, hList[j].n, distance)
			}
		}
	}
}

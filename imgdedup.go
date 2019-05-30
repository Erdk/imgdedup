package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var distanceFunction = flag.String("d", "manhattan", "distance type, default manhattan")
var toleration = flag.Uint("t", 100000, "distance under this value indicates that images are similar, default: 100000")
var dirpath = flag.String("p", "", "path to folder with images")
var verbose = flag.Bool("v", false, "turn on verbose messages")

func main() {
	flag.Parse()

	if *dirpath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir, err := filepath.Abs(*dirpath)
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
		fname := path.Join(*dirpath, f.Name())
		if *verbose {
			fmt.Printf("Processing (%v/%v): %v\n", i, numFiles, filepath.FromSlash(fname))
		} else {
			fmt.Printf("\r%d/%d", i, numFiles)
		}

		hist, err := histogram(fname)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			hList = append(hList, hist)
		}
	}

	fmt.Println()
	distFunc := getDistanceFunction(*distanceFunction)
	for i := 0; i < len(hList)-1; i++ {
		for j := i + 1; j < len(hList); j++ {
			distance := distFunc(hList[i].h, hList[j].h)
			if distance < *toleration {
				fmt.Printf("Similar: %v <-> %v  dist: %v\n",
					filepath.FromSlash(hList[i].n), filepath.FromSlash(hList[j].n), distance)
			}
		}
	}
}

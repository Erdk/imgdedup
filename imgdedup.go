package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/profile"
)

var distanceFunction = flag.String("t", "manhattan", "distance type, default manhattan")
var toleration = flag.Int("e", 100000, "distance under this value indicates that images are similar, default: 100000")
var dirpath = flag.String("d", "", "directory with images")
var verbose = flag.Bool("v", false, "turn on verbose messages")
var prof = flag.Bool("p", false, "profiling")

func main() {
	flag.Parse()

	if *dirpath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *prof {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}
	/*
		switch prof {
		case "cpu":
			defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
		case "mem":
			defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
		case "block":
			defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
		default:
		}
	*/

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
			fmt.Printf("\r%s\n", err.Error())
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

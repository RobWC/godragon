package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"sync"

	"github.com/robwc/godragon"
)

var currentVersion = "5.24.2"

var width = flag.Int("width", 1920, "Width of wallpapers")
var height = flag.Int("height", 1080, "Height of wallpapers")
var champname = flag.String("champ", "", "Specify a single champ to create wallpapers for")
var output = flag.String("output", ".", "Specify the output location for the wallpaper")

var wg sync.WaitGroup

func init() {
	flag.Parse()

	cn := *champname

	if len(cn) > 1 {
		fc := strings.ToUpper(string(cn[0]))
		rc := string(cn[1:len(cn)])
		cn = fmt.Sprintf("%s%s", fc, rc)
	} else if len(cn) == 1 {
		fc := strings.ToUpper(string(cn[0]))
		cn = fmt.Sprintf("%s", fc)
	}

	*champname = cn
}

func outputWallpaper(img image.Image, champName, pathLoc string) {
	endFile, err := os.Create(fmt.Sprintf("%s/Wallpaper-%s-%dx%d.png", pathLoc, champName, *width, *height))
	if err != nil {
		fmt.Printf("Error creating wallpaper : %s\n", err)
		os.Exit(1)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, img)
	wg.Done()
}

func main() {

	if *champname != "" {
		fmt.Printf("Creating wallpaper for %s at %dx%d\n", *champname, *width, *height)
		img, err := godragon.CreateWallpaper(*champname, currentVersion, *width, *height)
		if err != nil {
			fmt.Printf("Error creating wallpaper : %s\n", err)
			wg.Wait()
			os.Exit(1)
		}
		wg.Add(1)
		go outputWallpaper(img, *champname, *output)
	} else {
		champs, err := godragon.StaticChampions(currentVersion)
		if err != nil {
			fmt.Printf("Error fetching champion data : %s\n", err)
			os.Exit(1)
		}
		for k := range champs {
			fmt.Printf("Creating wallpaper for %s at %dx%d\n", k, *width, *height)
			img, err := godragon.CreateWallpaper(k, currentVersion, *width, *height)
			if err != nil {
				fmt.Printf("Error creating wallpaper : %s\n", err)
				wg.Wait()
				os.Exit(1)
			}
			wg.Add(1)
			go outputWallpaper(img, k, *output)
		}
	}
	wg.Wait()
	os.Exit(0)
}

package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/robwc/godragon"
)

var currentVersion = "5.24.2"

var width = flag.Int("width", 1920, "Width of wallpapers")
var height = flag.Int("height", 1080, "Height of wallpapers")
var champname = flag.String("champ", "", "Specify a single or multiple champion names to create a wallpaper (-champ Teemo or -champ Teemo,Ziggs)")
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

func main() {
	newchamp := "Teemo"

	c, err := godragon.StaticChampion(newchamp, currentVersion)
	if err != nil {
		fmt.Printf("Error getting infor for champion %s, ignoring champ\n", newchamp)
		os.Exit(1)
	}

	img, err := c.Image.Fetch(currentVersion)
	if err != nil {
		fmt.Printf("Error fetching image for champion %s, ignoring champ\n", newchamp)
		os.Exit(1)
	}

	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y

	// create filler image for blank spaces
	blackFill := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(blackFill, blackFill.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0}}, image.ZP, draw.Src)

	// calculate the size of the image
	// calculate the width of the image
	newImgWidthPx := len(newchamp) * (8 * imgWidth)
	newImgHeightPx := len(newchamp) * (8 * imgHeight)

	m := image.NewRGBA(image.Rect(0, 0, newImgWidthPx, newImgHeightPx))
	widthStartPx, heightStartPx := 0, 0

	for l := range newchamp {

		if v, ok := Letters[string(newchamp[l])]; !ok {
			log.Printf("Letter %s not available\n", string(newchamp[l]))
			os.Exit(1)
		} else {

			for i := 0; len(v) > i; i++ {
				for j := 0; len(v[i]) > j; j++ {
					log.Println(widthStartPx, heightStartPx, i, j, LetterA[i][j])
					if v[i][j] == 1 {
						draw.Draw(m, m.Bounds(), img, image.Point{X: widthStartPx, Y: heightStartPx}, draw.Src)
					}
					widthStartPx = widthStartPx - img.Bounds().Max.X

					if widthStartPx < -(newImgWidthPx) {
						break
					}

				}

				heightStartPx = heightStartPx - img.Bounds().Max.Y
				widthStartPx = 0
			}

		}
		widthStartPx = 0
		heightStartPx = heightStartPx + -((8 * imgHeight) + imgHeight)

	}
	godragon.WriteWallpaperFile(m, newchamp, *output)

}

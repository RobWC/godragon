package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
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
var all = flag.Bool("all", false, "Create wallpapers for all champs at the specified resolution")

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

	// split champnames
	champList := strings.Split(*champname, ",")

	if *all {
		fmt.Printf("Creating wallpapers for all champions at %dx%d\n", *width, *height)

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
			go func(name string) {
				godragon.WriteWallpaperFile(img, name, *output)
				wg.Done()
			}(k)
		}
	} else if len(champList) > 1 {

		type Champ struct {
			Name     string
			ID       int
			Champion godragon.Champion
			Img      image.Image
		}

		var champs []Champ

		for i := range champList {

			c, err := godragon.StaticChampion(champList[i], currentVersion)
			if err != nil {
				fmt.Printf("Error getting infor for champion %s, ignoring champ\n", champList[i])
				continue
			}

			img, err := c.Image.Fetch(currentVersion)
			if err != nil {
				fmt.Printf("Error fetching image for champion %s, ignoring champ\n", champList[i])
				continue
			}

			nc := Champ{
				Name:     champList[i],
				ID:       i,
				Champion: c,
				Img:      img,
			}

			champs = append(champs, nc)
		}

		maxWidth := *width
		maxHeight := *height

		m := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
		widthStart, heightStart := 0, 0

		x := 0
		for {

			draw.Draw(m, m.Bounds(), champs[x].Img, image.Point{X: widthStart, Y: heightStart}, draw.Src)
			widthStart = widthStart - champs[x].Img.Bounds().Max.X

			if widthStart < -(maxWidth) {
				if (widthStart < -(maxWidth)) && (heightStart < -(maxHeight)) {
					break
				}
				heightStart = heightStart - champs[x].Img.Bounds().Max.Y
				widthStart = 0
			}

			x++
			// reset champ index
			if x == len(champs) {
				x = 0
			}

		}

		champFinalList := []string{}
		for y := range champs {
			champFinalList = append(champFinalList, champs[y].Name)
		}

		wg.Add(1)
		go func() {
			godragon.WriteWallpaperFile(m, strings.Join(champFinalList, ","), *output)
			wg.Done()
		}()
	} else if len(champList) == 1 && champList[0] != "" {
		fmt.Printf("Creating wallpaper for %s at %dx%d\n", *champname, *width, *height)
		img, err := godragon.CreateWallpaper(*champname, currentVersion, *width, *height)
		if err != nil {
			fmt.Printf("Error creating wallpaper : %s\n", err)
			wg.Wait()
			os.Exit(1)
		}
		wg.Add(1)
		go func() {
			godragon.WriteWallpaperFile(img, *champname, *output)
			wg.Done()
		}()
	}
	wg.Wait()
	os.Exit(0)
}

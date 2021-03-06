package godragon

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// WriteWallpaperFile write file to JPEG at specified location
func WriteWallpaperFile(img image.Image, champName, pathLoc string) {
	if pathLoc == "" {
		pathLoc = "."
	}

	endFile, err := os.Create(fmt.Sprintf("%s/Wallpaper-%s-%dx%d.jpg", pathLoc, champName, img.Bounds().Max.X, img.Bounds().Max.Y))
	if err != nil {
		fmt.Printf("Error creating wallpaper : %s\n", err)
		os.Exit(1)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	jpeg.Encode(imgwr, img, &jpeg.Options{Quality: 100})
}

// CreateWallpaper create a tiled wallpaper of a specific champion
func CreateWallpaper(champ string, version string, width, height int) (image.Image, error) {
	c := Champion{}
	img, err := c.Image.Fetch(version)
	if err != nil {
		return nil, err
	}

	maxWidth := width
	maxHeight := height

	m := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
	widthStart, heightStart := 0, 0
	for {
		draw.Draw(m, m.Bounds(), img, image.Point{X: widthStart, Y: heightStart}, draw.Src)
		widthStart = widthStart - img.Bounds().Max.X
		if widthStart < -(maxWidth) {
			if (widthStart < -(maxWidth)) && (heightStart < -(maxHeight)) {
				break
			}
			heightStart = heightStart - img.Bounds().Max.Y
			widthStart = 0
		}
	}

	return m, nil
}

func CreateTestWallpaper(champName, version string, width, height int) error {
	c, err := StaticChampion(champName, version)
	if err != nil {
		return err
	}
	img, err := c.Image.Fetch(version)
	if err != nil {
		return err
	}

	maxWidth := width
	maxHeight := height

	m := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
	widthStart, heightStart := 0, 0
	for {
		draw.Draw(m, m.Bounds(), img, image.Point{X: widthStart, Y: heightStart}, draw.Src)
		widthStart = widthStart - img.Bounds().Max.X
		if widthStart < -(maxWidth) {
			if (widthStart < -(maxWidth)) && (heightStart < -(maxHeight)) {
				break
			}
			heightStart = heightStart - img.Bounds().Max.Y
			widthStart = 0
		}
	}
	_, err = os.Stat(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%s", champName))
	if os.IsNotExist(err) {
		// create directory
		err := os.Mkdir(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%s", champName), os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	endFile, err := os.Create(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%[1]s/Wallpaper-%[1]s-%dx%d.png", champName, width, height))
	if err != nil {
		return err
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, m)
	return nil
}

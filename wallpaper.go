package godragon

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// CreateWallpaper create a tiled wallpaper of a specific champion
func CreateWallpaper(champName, version string, width, height int) (image.Image, error) {
	c, err := StaticChampion(champName, version)
	if err != nil {
		return nil, err
	}
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

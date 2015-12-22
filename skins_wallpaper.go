package main

import (
	"fmt"
	"image"
	"image/draw"
)

// CreateSkinsWallpaper create a wallpaper from all skins of a champion
func CreateSkinsWallpaper(champName, version string, width, height int) (image.Image, error) {
	c, err := StaticChampion(champName, version)
	if err != nil {
		return nil, err
	}

	skinCount := len(c.Skins)

	if skinCount > 0 {
		maxWidth := width
		maxHeight := height

		// calculate on how we want to place the skins
		img, err := FetchChampLoadingImage(c.Name, c.Skins[0].Num)
		if err != nil {
			return nil, err
		}

		skinW := img.Bounds().Max.X
		skinH := img.Bounds().Max.Y
		skinMidH := skinH / 2
		skinTotalW := skinW * skinCount

		skinStart := 0
		if (maxWidth - skinTotalW) < 0 {
			// too many skins to fit so we need to fold to multiple lines
			return nil, fmt.Errorf("Too many skins to display correctly")
		} else {
			skinStart = (skinTotalW - maxWidth) / 2
		}

		wallpaperMidH := maxHeight / 2

		m := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
		widthStart, heightStart := 0, 0

		widthStart = skinStart

		for i := range c.Skins {
			img, err := FetchChampLoadingImage(c.Name, c.Skins[i].Num)
			if err != nil {
				return nil, err
			}

			draw.Draw(m, m.Bounds(), img, image.Point{X: widthStart, Y: skinMidH - wallpaperMidH}, draw.Src)
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
	return nil, fmt.Errorf("No skins found for champion %s", champName)
}

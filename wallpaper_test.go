package main

import (
	"bufio"
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestCreateWallpaper(t *testing.T) {
	_, err := CreateWallpaper("Teemo", "5.24.2", 1920, 1080)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateAllWallpaper(t *testing.T) {
	champs, err := StaticChampions(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	for k := range champs {
		_, err := CreateWallpaper(k, testVerion, 1920, 1080)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateSkinsWallpaper(t *testing.T) {
	champName := "Quinn"
	width := 1920
	height := 1080
	t.Logf("Creating Skins Wallpaper for %s at %dx%d", champName, width, height)
	img, err := CreateSkinsWallpaper(champName, testVerion, width, height)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%s", champName))
	if os.IsNotExist(err) {
		// create directory
		err := os.Mkdir(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%s", champName), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	} else if err != nil {
		t.Fatal(err)
	}

	endFile, err := os.Create(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%[1]s/SkinsWallpaper-%[1]s-%dx%d.png", champName, width, height))
	if err != nil {
		t.Fatal(err)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, img)
}

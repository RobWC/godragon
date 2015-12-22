package main

import (
	"bufio"
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestCreateWallpaper(t *testing.T) {
	champName := "Quinn"
	width := 1920
	height := 1080
	img, err := CreateWallpaper(champName, testVerion, width, height)
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

	endFile, err := os.Create(fmt.Sprintf("/Users/rcameron/gopath/src/github.com/robwc/godragon/test/%[1]s/Wallpaper-%[1]s-%dx%d.png", champName, width, height))
	if err != nil {
		t.Fatal(err)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, img)
}

func TestCreateAllWallpaper(t *testing.T) {
	width := 1920
	height := 1080

	champs, err := StaticChampions(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	for k := range champs {
		_, err := CreateWallpaper(k, testVerion, width, height)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateSkinsWallpaper(t *testing.T) {
	champName := "Braum"
	width := 2880
	height := 1800
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

package godragon

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestImageFetchChampion(t *testing.T) {
	c, err := StaticChampion("Braum", testVerion)
	if err != nil {
		t.Fatal(err)
	}
	full, err := c.Image.Fetch(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	fullFile, err := os.Create(fmt.Sprintf("test/TestImageFetchFull.png"))
	if err != nil {
		t.Fatal(err)
	}
	defer fullFile.Close()
	fullwr := bufio.NewWriter(fullFile)
	png.Encode(fullwr, full)
}

func TestFetchSprite(t *testing.T) {
	c, err := StaticChampion("Braum", testVerion)
	if err != nil {
		t.Fatal(err)
	}
	sprite, err := c.Image.FetchSprite(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	spriteFile, err := os.Create(fmt.Sprintf("test/TestImageFetchSprite.jpg"))
	if err != nil {
		t.Fatal(err)
	}
	defer spriteFile.Close()
	spritewr := bufio.NewWriter(spriteFile)
	err = jpeg.Encode(spritewr, sprite, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}

}

func TestFetchChampLoadingImage(t *testing.T) {

	img, err := FetchChampLoadingImage("Braum", 0)
	if err != nil {
		t.Fatal(err)
	}
	endFile, err := os.Create(fmt.Sprintf("test/TestImageFetchChampLoadingImage.png"))
	if err != nil {
		t.Fatal(err)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, img)
}

func TestFetchChampSplashImage(t *testing.T) {
	img, err := FetchChampSplashImage("Braum", 0)
	if err != nil {
		t.Fatal(err)
	}
	endFile, err := os.Create(fmt.Sprintf("test/TestImageFetchChampSplashImage.png"))
	if err != nil {
		t.Fatal(err)
	}
	defer endFile.Close()
	imgwr := bufio.NewWriter(endFile)
	png.Encode(imgwr, img)
}

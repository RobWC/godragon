package main

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestImageFetch(t *testing.T) {
	c, err := StaticChampion("5.14.1", "Annie")
	if err != nil {
		t.Fatal(err)
	}
	img, err := c.Image.FetchFull("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile()
	f, err := os.OpenFile("./test/TestImageFetchFull.png", os.O_RDWR, 666)
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}

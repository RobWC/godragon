package main

import (
	"io/ioutil"
	"testing"
)

func TestImageFetchChampion(t *testing.T) {
	c, err := StaticChampion(testVerion, "Braum")
	if err != nil {
		t.Fatal(err)
	}
	full, err := c.Image.FetchFull(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test/TestImageFetchFull.png", full, 0777)
	if err != nil {
		t.Fatal(err)
	}
	sprite, err := c.Image.FetchSprite(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test/TestImageFetchSprite.png", sprite, 0777)
	if err != nil {
		t.Fatal(err)
	}
	data, err := FetchChampLoadingImage("Braum", 0)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test/TestImageFetchChampLoadingImage.png", data, 0777)
	if err != nil {
		t.Fatal(err)
	}
}

package main

import (
	"io/ioutil"
	"testing"
)

func TestImageFetchChampion(t *testing.T) {
	c, err := StaticChampion("5.14.1", "Annie")
	if err != nil {
		t.Fatal(err)
	}
	full, err := c.Image.FetchFull("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test/TestImageFetchFull.png", full, 0777)
	if err != nil {
		t.Fatal(err)
	}
	sprite, err := c.Image.FetchSprite("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test/TestImageFetchSprite.png", sprite, 0777)
	if err != nil {
		t.Fatal(err)
	}
}

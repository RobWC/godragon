package main

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
)

type Image struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

func (i *Image) FetchSprite(version string) (sprite image.Image, err error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/sprite/%s", version, i.Sprite))
	if err != nil {
		return sprite, err
	}
	if resp.Header.Get("Content-Type") == "image/png" {
		sprite, err = png.Decode(resp.Body)
		if err != nil {
			return sprite, err
		}
	}
	return sprite, err
}

func (i *Image) FetchFull(version string) (full image.Image, err error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/%s/%s", version, i.Group, i.Sprite))
	if err != nil {
		return full, err
	}
	if resp.Header.Get("Content-Type") == "image/png" {
		full, err = png.Decode(resp.Body)
		if err != nil {
			return full, err
		}
	}
	return full, err
}

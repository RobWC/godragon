package main

import (
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"
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

func (i *Image) FetchSprite(version string) (sprite []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/sprite/%s", version, i.Sprite))
	if err != nil {
		return sprite, err
	}
	if resp.Header.Get("Content-Type") == "image/png" {
		sprite, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return sprite, err
		}
	}
	return sprite, err
}

func (i *Image) FetchFull(version string) (full []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/%s/%s", version, i.Group, i.Full))
	if err != nil {
		return full, err
	}
	if resp.Header.Get("Content-Type") == "image/png" {
		full, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return full, err
		}
	}
	return full, err
}

// FetchChampLoadingImage fetch the loading image for a Champion with specified skin
func FetchChampLoadingImage(n string, s int) (data []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/img/champion/loading/%s_%s.jpg", strings.Join([]string{strings.ToUpper(string(n[0])), strings.ToLower(string(n[1:len(n)]))}, ""), strconv.Itoa(s)))
	if err != nil {
		return data, err
	}
	if resp.StatusCode == 200 {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return data, err
		}
	}
	return data, err
}

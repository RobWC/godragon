package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"strconv"
	"strings"

	"io/ioutil"
	"net/http"
)

// Image image construct for static images
type Image struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

// FetchSprite return the sprite of the image
func (i *Image) FetchSprite(version string) (image.Image, error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/sprite/%s", version, i.Sprite))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if (resp.StatusCode == 200) && (resp.Header.Get("Content-Type") == "image/png") {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		img, err := png.Decode(buf)
		if err != nil {
			return nil, err
		}
		log.Println("coord", i.X, i.Y, i.W, i.H)
		m := image.NewRGBA(image.Rect(0, 0, i.W, i.H))
		draw.Draw(m, m.Bounds(), img, image.Point{X: i.X, Y: i.Y}, draw.Src)

		return m, err
	}
	return nil, fmt.Errorf("Error fetching image %s response code %d", i.Full, resp.StatusCode)
}

// Fetch fetch the full image
func (i *Image) Fetch(version string) (image.Image, error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/%s/%s", version, i.Group, i.Full))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if (resp.StatusCode == 200) && (resp.Header.Get("Content-Type") == "image/png") {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		img, err := png.Decode(buf)
		if err != nil {
			return nil, err
		}
		return img, err
	}
	return nil, fmt.Errorf("Error fetching image %s response code %d", i.Full, resp.StatusCode)
}

// FetchChampLoadingImage fetch the loading image for a Champion with specified skin
func FetchChampLoadingImage(n string, s int) (image.Image, error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/img/champion/loading/%s_%s.jpg", strings.Join([]string{strings.ToUpper(string(n[0])), strings.ToLower(string(n[1:len(n)]))}, ""), strconv.Itoa(s)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		img, err := jpeg.Decode(buf)
		if err != nil {
			return nil, err
		}
		return img, err
	}
	return nil, fmt.Errorf("Error fetching image for champ %s and skin %d response code %d", n, s, resp.StatusCode)
}

// FetchChampSplashImage fetch the splash image for a Champion with specified skin
func FetchChampSplashImage(n string, s int) (image.Image, error) {
	resp, err := http.Get(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/img/champion/splash/%s_%s.jpg", strings.Join([]string{strings.ToUpper(string(n[0])), strings.ToLower(string(n[1:len(n)]))}, ""), strconv.Itoa(s)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		img, err := jpeg.Decode(buf)
		if err != nil {
			return nil, err
		}
		return img, err
	}
	return nil, fmt.Errorf("Error fetching image for champ %s and skin %d response code %d", n, s, resp.StatusCode)
}

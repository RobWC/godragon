package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StaticChampions(version string) (cr map[string]Champion, err error) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", version)
	resp, err := http.Get(path)
	if err != nil {
		return cr, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cr, err
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	return cd.Champions, err
}

func StaticChampion(name string, version string) (c Champion, err error) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion/%s.json", version, name)
	resp, err := http.Get(path)
	if err != nil {
		return c, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c, err
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	return cd.Champions[name], err
}

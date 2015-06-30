package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ChampionData struct {
	Type      string              `json:"type"`
	Format    string              `json:"format"`
	Version   string              `json:"version"`
	Champions map[string]Champion `json:"data"`
}

type Champion struct {
	Version string        `json:"version"`
	ID      string        `json:"id"`
	Key     string        `json:"key"`
	Name    string        `json:"name"`
	Title   string        `json:"title"`
	Blurb   string        `json:"blurb"`
	Info    ChampionInfo  `json:"info"`
	Image   ChampionImage `json:"image"`
	Tags    []string      `json:"tags"`
	ParType string        `json:"partype"`
	Stats   ChampionStats `json:"stats"`
}

type ChampionInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

type ChampionImage struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

type ChampionStats struct {
	HP                   float32 `json:"hp"`
	HPPerLevel           float32 `json:"hpperlevel"`
	MP                   float32 `json:"mp"`
	MPPerLevel           float32 `json:"mpperlevel"`
	MoveSpeed            float32 `json:"movespeed"`
	Armor                float32 `json:"armor"`
	ArmorPerLevel        float32 `json:"armorperlevel"`
	SpellBlock           float32 `json:"spellblock"`
	SpellBlockPerLevel   float32 `json:"spellblockperlevel"`
	AttackRange          float32 `json:"attackrange"`
	HpRegen              float32 `json:"hpregen"`
	HPRegenPerLevel      float32 `json:"hpregenperlevel"`
	MPRegen              float32 `json:"mpregen"`
	MPRegenPerLevel      float32 `json:"mpregenperlevel"`
	Crit                 float32 `json:"crit"`
	CritPerLevel         float32 `json:"critperlevel"`
	AttackDamage         float32 `json:"attackdamage"`
	AttackDamagePerLevel float32 `json:"attackdamageperlevel"`
	AttackSpeedOffset    float32 `json:"attackspeedoffset"`
	AttackSpeedPerLevel  float32 `json:"attackspeedperlevel"`
}

func StaticChampions(version string) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", version)
	resp, err := http.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	log.Printf("%#v", cd)
	log.Println("Total Champs:", len(cd.Champions))
}

type VersionList struct {
	Versions []string
}

func StaticVersions() {
	path := "https://ddragon.leagueoflegends.com/api/versions.json"
	resp, err := http.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	vl := &VersionList{}
	json.Unmarshal(body, &vl.Versions)
	log.Printf("%#v", vl.Versions)
	log.Println("Total Versions:", len(vl.Versions))
}

type LanguageList struct {
	Languages []string
}

func StaticLanguages() {
	path := "https://ddragon.leagueoflegends.com/cdn/languages.json"
	resp, err := http.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	ll := &LanguageList{}
	json.Unmarshal(body, &ll.Languages)
	log.Printf("%#v", ll.Languages)
	log.Println("Total Versions:", len(ll.Languages))
}

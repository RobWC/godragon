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
	Version   string          `json:"version"`
	ID        string          `json:"id"`
	Key       string          `json:"key"`
	Name      string          `json:"name"`
	Title     string          `json:"title"`
	Blurb     string          `json:"blurb"`
	AllyTips  []string        `json:"allytips"`
	EnemyTips []string        `json:"enemytips"`
	Spells    []ChampionSpell `json:"Spells"`
	Info      ChampionInfo    `json:"info"`
	Image     Image           `json:"image"`
	Skins     []ChampionSkin  `json:"skins"`
	Lore      string          `json:"lore"`
	Tags      []string        `json:"tags"`
	ParType   string          `json:"partype"`
	Stats     ChampionStats   `json:"stats"`
}

type ChampionSkin struct {
	ID   string `json:"id"`
	Num  int    `json:"num"`
	Name string `json:"name"`
}

type ChampionSpell struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	ToolTip      string              `json:"tooltip"`
	LevelTip     map[string][]string `json:"leveltip"`
	MaxRank      int                 `json:"maxrank"`
	Cooldown     []int               `json:"cooldown"`
	CooldownBurn string              `json:"cooldownBurn"`
	Cost         []int               `json:"cost"`
	CostBurn     string              `json:"costBurn"`
	Effect       []string            `json:"effect"` //TODO: What is this really?
	EffectBurn   []string            `json:"effectburn"`
	Vars         []SpellVar          `json:"vars"`
	CostType     string              `json:"costType"`
	Range        []int               `json:"range"`
	RangeBurn    string              `json:"rangeBurn"`
	Image        Image               `json:"image"`
	Resource     string              `json:"resource"`
}

type SpellVar struct {
	Link  string  `json:"link"`
	Coeff float32 `json:"coeff"`
	Key   string  `json:"key"`
}

type ChampionInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty" gorethink:"difficulty"`
}

type Image struct {
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

func StaticChampions(version string) (err error, cr map[string]Champion) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", version)
	resp, err := http.Get(path)
	if err != nil {
		return err, cr
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, cr
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	log.Printf("%#v", cd)
	log.Println("Total Champs:", len(cd.Champions))
	return err, cd.Champions
}

func StaticChampion(version string, name string) (err error, c Champion) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion/%s.json", version, name)
	resp, err := http.Get(path)
	if err != nil {
		return err, c
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, c
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	return err, cd.Champions[name]
}

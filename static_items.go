package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// http://ddragon.leagueoflegends.com/cdn/5.2.1/data/en_US/item.json

/**
"basic": {
        "name": "",
        "rune": {
            "isrune": false,
            "tier": 1,
            "type": "red"
        },
        "gold": {
            "base int
            "total int
            "sell int
            "purchasable": false
        },
        "group": "",
        "description": "",
        "colloq": ";",
        "plaintext": "",
        "consumed": false,
        "stacks": 1,
        "depth": 1,
        "consumeOnFull": false,
        "from": [],
        "into": [],
        "specialRecipe int
        "inStore": true,
        "hideFromAll": false,
        "requiredChampion": "",
        "stats": {
            "FlatHPPoolMod int
            "rFlatHPModPerLevel int
            "FlatMPPoolMod int
            "rFlatMPModPerLevel int
            "PercentHPPoolMod int
            "PercentMPPoolMod int
            "FlatHPRegenMod int
            "rFlatHPRegenModPerLevel int
            "PercentHPRegenMod int
            "FlatMPRegenMod int
            "rFlatMPRegenModPerLevel int
            "PercentMPRegenMod int
            "FlatArmorMod int
            "rFlatArmorModPerLevel int
            "PercentArmorMod int
            "rFlatArmorPenetrationMod int
            "rFlatArmorPenetrationModPerLevel int
            "rPercentArmorPenetrationMod int
            "rPercentArmorPenetrationModPerLevel int
            "FlatPhysicalDamageMod int
            "rFlatPhysicalDamageModPerLevel int
            "PercentPhysicalDamageMod int
            "FlatMagicDamageMod int
            "rFlatMagicDamageModPerLevel int
            "PercentMagicDamageMod int
            "FlatMovementSpeedMod int
            "rFlatMovementSpeedModPerLevel int
            "PercentMovementSpeedMod int
            "rPercentMovementSpeedModPerLevel int
            "FlatAttackSpeedMod int
            "PercentAttackSpeedMod int
            "rPercentAttackSpeedModPerLevel int
            "rFlatDodgeMod int
            "rFlatDodgeModPerLevel int
            "PercentDodgeMod int
            "FlatCritChanceMod int
            "rFlatCritChanceModPerLevel int
            "PercentCritChanceMod int
            "FlatCritDamageMod int
            "rFlatCritDamageModPerLevel int
            "PercentCritDamageMod int
            "FlatBlockMod int
            "PercentBlockMod int
            "FlatSpellBlockMod int
            "rFlatSpellBlockModPerLevel int
            "PercentSpellBlockMod int
            "FlatEXPBonus int
            "PercentEXPBonus int
            "rPercentCooldownMod int
            "rPercentCooldownModPerLevel int
            "rFlatTimeDeadMod int
            "rFlatTimeDeadModPerLevel int
            "rPercentTimeDeadMod int
            "rPercentTimeDeadModPerLevel int
            "rFlatGoldPer10Mod int
            "rFlatMagicPenetrationMod int
            "rFlatMagicPenetrationModPerLevel int
            "rPercentMagicPenetrationMod int
            "rPercentMagicPenetrationModPerLevel int
            "FlatEnergyRegenMod int
            "rFlatEnergyRegenModPerLevel int
            "FlatEnergyPoolMod int
            "rFlatEnergyModPerLevel int
            "PercentLifeStealMod int
            "PercentSpellVampMod": 0
        },
        "tags": [],
        "maps": {
            "1": true,
            "8": true,
            "10": true,
            "12": true
        }
    }
**/
type Item struct {
	Name        string     `json:"name"`
	Group       string     `json:"group"`
	Description string     `json:"description"`
	Colloq      string     `json:"colloq"`
	Plaintext   string     `json:"plaintext"`
	Into        []string   `json:"into"`
	Image       Image      `json:"image"`
	Cost        ItemCost   `json:"gold"`
	Tags        []string   `json:"tags"`
	Stats       ItemStats  `json:"stats"`
	Effect      ItemEffect `json:"effect"`
	Maps        MapList    `json:"maps"`
}

type MapList struct {
	Map1  bool `json:"1"`
	Map8  bool `json:"8"`
	Map10 bool `json:"10"`
	Map12 bool `json:"12"`
}

type ItemCost struct {
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
	Total       int  `json:"total"`
	Sell        int  `json:"sell"`
}

type ItemStats struct {
	FlatHPPoolMod                       int `json:"FlatHPPoolMod"`
	rFlatHPModPerLevel                  int `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       int `json:"FlatMPPoolMod"`
	rFlatMPModPerLevel                  int `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    int `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    int `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      int `json:"FlatHPRegenMod"`
	rFlatHPRegenModPerLevel             int `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   int `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      int `json:"FlatMPRegenMod"`
	rFlatMPRegenModPerLevel             int `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   int `json:"PercentMPRegenMod"`
	FlatArmorMod                        int `json:"FlatArmorMod"`
	rFlatArmorModPerLevel               int `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     int `json:"PercentArmorMod"`
	rFlatArmorPenetrationMod            int `json:"rFlatArmorPenetrationMod"`
	rFlatArmorPenetrationModPerLevel    int `json:"rFlatArmorPenetrationModPerLevel"`
	rPercentArmorPenetrationMod         int `json:"rPercentArmorPenetrationMod"`
	rPercentArmorPenetrationModPerLevel int `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               int `json:"FlatPhysicalDamageMod"`
	rFlatPhysicalDamageModPerLevel      int `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            int `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  int `json:"FlatMagicDamageMod"`
	rFlatMagicDamageModPerLevel         int `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               int `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                int `json:"FlatMovementSpeedMod"`
	rFlatMovementSpeedModPerLevel       int `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             int `json:"PercentMovementSpeedMod"`
	rPercentMovementSpeedModPerLevel    int `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  int `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               int `json:"PercentAttackSpeedMod"`
	rPercentAttackSpeedModPerLevel      int `json:"rPercentAttackSpeedModPerLevel"`
	rFlatDodgeMod                       int `json:"rFlatDodgeMod"`
	rFlatDodgeModPerLevel               int `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     int `json:"PercentDodgeMod"`
	FlatCritChanceMod                   int `json:"FlatCritChanceMod"`
	rFlatCritChanceModPerLevel          int `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                int `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   int `json:"FlatCritDamageMod"`
	rFlatCritDamageModPerLevel          int `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                int `json:"PercentCritDamageMod"`
	FlatBlockMod                        int `json:"FlatBlockMod"`
	PercentBlockMod                     int `json:"PercentBlockMod"`
	FlatSpellBlockMod                   int `json:"FlatSpellBlockMod"`
	rFlatSpellBlockModPerLevel          int `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                int `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        int `json:"FlatEXPBonus"`
	PercentEXPBonus                     int `json:"PercentEXPBonus"`
	rPercentCooldownMod                 int `json:"rPercentCooldownMod"`
	rPercentCooldownModPerLevel         int `json:"rPercentCooldownModPerLevel"`
	rFlatTimeDeadMod                    int `json:"rFlatTimeDeadMod"`
	rFlatTimeDeadModPerLevel            int `json:"rFlatTimeDeadModPerLevel"`
	rPercentTimeDeadMod                 int `json:"rPercentTimeDeadMod"`
	rPercentTimeDeadModPerLevel         int `json:"rPercentTimeDeadModPerLevel"`
	rFlatGoldPer10Mod                   int `json:"rFlatGoldPer10Mod"`
	rFlatMagicPenetrationMod            int `json:"rFlatMagicPenetrationMod"`
	rFlatMagicPenetrationModPerLevel    int `json:"rFlatMagicPenetrationModPerLevel"`
	rPercentMagicPenetrationMod         int `json:"rPercentMagicPenetrationMod"`
	rPercentMagicPenetrationModPerLevel int `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  int `json:"FlatEnergyRegenMod"`
	rFlatEnergyRegenModPerLevel         int `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   int `json:"FlatEnergyPoolMod"`
	rFlatEnergyModPerLevel              int `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 int `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 int `json:"PercentSpellVampMod"`
}

type ItemEffect struct {
	Effect1Amount string `json:"Effect1Amount"`
	Effect2Amount string `json:"Effect2Amount"`
	Effect3Amount string `json:"Effect3Amount"`
	Effect4Amount string `json:"Effect4Amount"`
	Effect5Amount string `json:"Effect5Amount"`
	Effect6Amount string `json:"Effect6Amount"`
	Effect7Amount string `json:"Effect7Amount"`
	Effect8Amount string `json:"Effect8Amount"`
}

type ItemData struct {
	Type    string          `json:"type"`
	Format  string          `json:"format"`
	Version string          `json:"version"`
	Basic   string          `json:"-"`
	Items   map[string]Item `json:"data"`
}

func StaticItems(version string) (err error, ir map[string]Item) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/item.json", version)
	resp, err := http.Get(path)
	if err != nil {
		return err, ir
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ir
	}
	id := ItemData{}
	json.Unmarshal(body, &id)
	log.Printf("%#v", id)
	log.Println("Total Items:", len(id.Items))
	return err, id.Items
}

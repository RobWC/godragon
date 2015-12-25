package godragon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	ID          string
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
	FlatHPPoolMod                       float32 `json:"FlatHPPoolMod"`
	rFlatHPModPerLevel                  float32 `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       float32 `json:"FlatMPPoolMod"`
	rFlatMPModPerLevel                  float32 `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    float32 `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    float32 `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      float32 `json:"FlatHPRegenMod"`
	rFlatHPRegenModPerLevel             float32 `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   float32 `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      float32 `json:"FlatMPRegenMod"`
	rFlatMPRegenModPerLevel             float32 `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   float32 `json:"PercentMPRegenMod"`
	FlatArmorMod                        float32 `json:"FlatArmorMod"`
	rFlatArmorModPerLevel               float32 `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     float32 `json:"PercentArmorMod"`
	rFlatArmorPenetrationMod            float32 `json:"rFlatArmorPenetrationMod"`
	rFlatArmorPenetrationModPerLevel    float32 `json:"rFlatArmorPenetrationModPerLevel"`
	rPercentArmorPenetrationMod         float32 `json:"rPercentArmorPenetrationMod"`
	rPercentArmorPenetrationModPerLevel float32 `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               float32 `json:"FlatPhysicalDamageMod"`
	rFlatPhysicalDamageModPerLevel      float32 `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            float32 `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  float32 `json:"FlatMagicDamageMod"`
	rFlatMagicDamageModPerLevel         float32 `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               float32 `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                float32 `json:"FlatMovementSpeedMod"`
	rFlatMovementSpeedModPerLevel       float32 `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             float32 `json:"PercentMovementSpeedMod"`
	rPercentMovementSpeedModPerLevel    float32 `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  float32 `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               float32 `json:"PercentAttackSpeedMod"`
	rPercentAttackSpeedModPerLevel      float32 `json:"rPercentAttackSpeedModPerLevel"`
	rFlatDodgeMod                       float32 `json:"rFlatDodgeMod"`
	rFlatDodgeModPerLevel               float32 `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     float32 `json:"PercentDodgeMod"`
	FlatCritChanceMod                   float32 `json:"FlatCritChanceMod"`
	rFlatCritChanceModPerLevel          float32 `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                float32 `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   float32 `json:"FlatCritDamageMod"`
	rFlatCritDamageModPerLevel          float32 `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                float32 `json:"PercentCritDamageMod"`
	FlatBlockMod                        float32 `json:"FlatBlockMod"`
	PercentBlockMod                     float32 `json:"PercentBlockMod"`
	FlatSpellBlockMod                   float32 `json:"FlatSpellBlockMod"`
	rFlatSpellBlockModPerLevel          float32 `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                float32 `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        float32 `json:"FlatEXPBonus"`
	PercentEXPBonus                     float32 `json:"PercentEXPBonus"`
	rPercentCooldownMod                 float32 `json:"rPercentCooldownMod"`
	rPercentCooldownModPerLevel         float32 `json:"rPercentCooldownModPerLevel"`
	rFlatTimeDeadMod                    float32 `json:"rFlatTimeDeadMod"`
	rFlatTimeDeadModPerLevel            float32 `json:"rFlatTimeDeadModPerLevel"`
	rPercentTimeDeadMod                 float32 `json:"rPercentTimeDeadMod"`
	rPercentTimeDeadModPerLevel         float32 `json:"rPercentTimeDeadModPerLevel"`
	rFlatGoldPer10Mod                   float32 `json:"rFlatGoldPer10Mod"`
	rFlatMagicPenetrationMod            float32 `json:"rFlatMagicPenetrationMod"`
	rFlatMagicPenetrationModPerLevel    float32 `json:"rFlatMagicPenetrationModPerLevel"`
	rPercentMagicPenetrationMod         float32 `json:"rPercentMagicPenetrationMod"`
	rPercentMagicPenetrationModPerLevel float32 `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  float32 `json:"FlatEnergyRegenMod"`
	rFlatEnergyRegenModPerLevel         float32 `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   float32 `json:"FlatEnergyPoolMod"`
	rFlatEnergyModPerLevel              float32 `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 float32 `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 float32 `json:"PercentSpellVampMod"`
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

func StaticItems(version string) (ir map[string]Item, err error) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/item.json", version)
	resp, err := http.Get(path)
	if err != nil {
		return ir, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ir, err
	}
	id := ItemData{}
	json.Unmarshal(body, &id)
	for item := range id.Items {
		updateItem := id.Items[item]
		updateItem.ID = item
		id.Items[item] = updateItem
	}
	return id.Items, err
}

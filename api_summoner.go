package godragon

import (
	"encoding/json"
	"strings"
)

type Account struct {
	AccountID int
	PlatforID string
}

type Summoner struct {
	AccountID    int            `json:"accountId"`
	SummonerID   int            `json:"summonerId"`
	PlatformID   string         `json:"platformId"`
	AllyBadge    int            `json:"allyBadge"`
	CurrentUser  Account        `json:"currentUser"`
	ELO          map[string]int `json:"elo"`
	EnemeyBadge  int            `json:"enemyBadge"`
	HighestRank  string         `json:"highestRank"`
	SummonerName string         `json:"summonerName"`
	ProfileIcon  int            `json:"profileIcon"`
}

type SummonerList map[string]*Summoner

// SummonerByName v2.2 of get summoner by name
func (c *APIClient) SummonerByName(name string) (*Summoner, error) {

	apiString := strings.Join([]string{"summoner", "by-name", name}, "/")
	req, err := c.genRequest("GET", "v2.2", apiString, nil)
	if err != nil {
		return nil, err
	}
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	sl := make(SummonerList)
	err = json.Unmarshal(data, &sl)
	if err != nil {
		return nil, err
	}

	return sl[name], nil
}

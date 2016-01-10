package godragon

import (
	"os"
	"testing"
)

func TestChampionStatsBySummoner(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOT_API_KEY")

	c := NewAPIClient("na", apiKeyEnv)
	cs, err := c.ChampionStatsBySummoner(46779953)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cs)
}

package godragon

import (
	"os"
	"testing"
)

func TestSummonerByName(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOT_API_KEY")

	c := NewAPIClient("na", apiKeyEnv)
	s, err := c.SummonerByName("MrCrapper")
	if err != nil {
		t.Fatal(s)
	}
	t.Log(s)
}

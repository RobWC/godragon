package godragon

import "testing"

func TestChampionCooldowns(t *testing.T) {

	nc, err := StaticChampion("Sivir", "6.4.1")
	if err != nil {
		t.Fatal(err)
	}

	err = ChampionCooldowns(nc)
	if err != nil {
		t.Fatal(err)
	}
}

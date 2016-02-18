package godragon

import "testing"

func TestChampionCooldowns(t *testing.T) {

	nc, err := StaticChampion("Nami", "6.3.1")
	if err != nil {
		t.Fatal(err)
	}

	err = ChampionCooldowns(nc)
	if err != nil {
		t.Fatal(err)
	}
}

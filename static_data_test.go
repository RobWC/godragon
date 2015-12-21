package main

import "testing"

var testVerion = "5.24.1"

func TestStaticChampions(t *testing.T) {
	_, err := StaticChampions(testVerion)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticChampion(t *testing.T) {
	_, err := StaticChampion(testVerion, "Annie")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticVersions(t *testing.T) {
	StaticVersions()
}

func TestStaticLanguages(t *testing.T) {
	StaticLanguages()
}

func TestChampionList(t *testing.T) {
	champs, err := StaticChampions(testVerion)
	if err != nil {
		t.Fatal(err)
	}
	for k := range champs {
		_, err := StaticChampion(testVerion, k)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestItemList(t *testing.T) {
	_, err := StaticItems(testVerion)
	if err != nil {
		t.Fatal(err)
	}
}

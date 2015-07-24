package main

import "testing"

func TestStaticChampions(t *testing.T) {
	_, err := StaticChampions("5.12.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticChampion(t *testing.T) {
	_, err := StaticChampion("5.12.1", "Annie")
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
	champs, err := StaticChampions("5.12.1")
	if err != nil {
		t.Fatal(err)
	}
	for k := range champs {
		_, err := StaticChampion("5.12.1", k)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestItemList(t *testing.T) {
	_, err := StaticItems("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
}

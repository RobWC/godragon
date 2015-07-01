package main

import "testing"

func TestStaticChampions(t *testing.T) {
	StaticChampions("5.12.1")
}

func TestStaticChampion(t *testing.T) {
	StaticChampion("5.12.1", "Annie")
}

func TestStaticVersions(t *testing.T) {
	StaticVersions()
}

func TestStaticLanguages(t *testing.T) {
	StaticLanguages()
}

func TestChampionList(t *testing.T) {
	err, champs := StaticChampions("5.12.1")
	if err != nil {
		t.Fatal(err)
	}
	for k := range champs {
		err, _ := StaticChampion("5.12.1", k)
		if err != nil {
			t.Fatal(err)
		}
	}
}

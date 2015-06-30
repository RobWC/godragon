package main

import "testing"

func TestStaticChampions(t *testing.T) {
	StaticChampions("5.12.1")
}

func TestStaticVersions(t *testing.T) {
	StaticVersions()
}

func TestStaticLanguages(t *testing.T) {
	StaticLanguages()
}

package godragon

import (
	"log"
	"os"
	"testing"
)

func TestNewStaticChampions(t *testing.T) {
	cd, err := NewStaticChampions(os.Getenv("RIOT_KEY"))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(cd)

}

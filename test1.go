package main

import (
	"log"
	"os"
)

func main() {
	items, err := StaticItems("5.24.2")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for k, v := range items {
		log.Println(k, v.Name, v.Stats.FlatPhysicalDamageMod)
	}
}

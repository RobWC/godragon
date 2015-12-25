package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/robwc/godragon"
)

func main() {
	items, err := godragon.StaticItems("5.24.2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session, err := mgo.Dial("192.168.99.100")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer session.Close()

	for _, v := range items {
		c := session.DB("lol").C("items")

		err = c.Insert(v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

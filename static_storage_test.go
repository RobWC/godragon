package main

import (
	"log"
	"testing"
	"time"
)

func TestStaticStorageConn(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticStorageSchema(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = s.CreateSchema()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticStorageDropSchema(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = s.DropSchema()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticStorageInsert(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = s.CreateSchema()
	if err != nil {
		t.Fatal(err)
	}
	var c Champion
	err, c = StaticChampion("5.12.1", "Annie")
	if err != nil {
		t.Fatal(err)
	}
	err = s.AddChampion(c)
	if err != nil {
		t.Fatal(err)
	}
	err = s.DropSchema()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticStorageInsertAll(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = s.CreateSchema()
	if err != nil {
		t.Fatal(err)
	}
	c := make(map[string]Champion)
	err, c = StaticChampions("5.12.1")
	if err != nil {
		t.Fatal(err)
	}
	for k := range c {
		log.Println("Fetching", k)
		err, champ := StaticChampion("5.12.1", k)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)
		log.Println("Inserting", k)
		err = s.AddChampion(champ)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}

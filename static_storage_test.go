package main

import (
	"log"
	"testing"
)

func TestStaticStorageConn(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
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
		log.Println(err)
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
	c, err = StaticChampion("5.12.1", "Annie")
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
	c, err = StaticChampions("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
	for k := range c {
		champ, err := StaticChampion("5.14.1", k)
		if err != nil {
			t.Fatal(err)
		}
		err = s.AddChampion(champ)
		if err != nil {
			t.Fatal(err)
		}
	}
	items, err := StaticItems("5.14.1")
	if err != nil {
		t.Fatal(err)
	}
	for k := range items {
		err = s.AddItem(items[k])
		if err != nil {
			t.Fatal(err)
		}
	}
	//err = s.DropSchema()
	//if err != nil {
	//	t.Fatal(err)
	//}
}

package main

import "testing"

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

func TestStaticStorageInsert(t *testing.T) {
	s := &StaticStorage{}
	err := s.Connect()
	if err != nil {
		t.Fatal(err)
	}
	/**
	err = s.CreateSchema()
	if err != nil {
		t.Fatal(err)
	}**/
	c := StaticChampion("5.12.1", "Annie")
	err = s.AddChampion(c)
	if err != nil {
		t.Fatal(err)
	}
}

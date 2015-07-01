package main

import (
	"time"

	r "github.com/dancannon/gorethink"
)

type StaticStorage struct {
	Session *r.Session
}

func (s *StaticStorage) Connect() (err error) {
	s.Session, err = r.Connect(
		r.ConnectOpts{Address: "localhost:28015",
			Database:      "lol",
			Timeout:       5 * time.Second,
			MaxIdle:       100,
			MaxOpen:       1000,
			DiscoverHosts: true})
	return err
}

func (s *StaticStorage) CreateSchema() error {
	_, err := r.DB("lol").TableCreate("champions").RunWrite(s.Session)
	if err != nil {
		return err
	}
	return nil
}

func (s *StaticStorage) DropSchema() error {
	_, err := r.DB("lol").TableDrop("champions").RunWrite(s.Session)
	if err != nil {
		return err
	}
	return nil
}

func (s *StaticStorage) AddChampion(champ Champion) (err error) {
	_, err = r.Table("champions").Insert(champ).RunWrite(s.Session)
	if err != nil {
		return err
	}
	return nil
}

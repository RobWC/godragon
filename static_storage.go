package main

import r "github.com/dancannon/gorethink"

type StaticStorage struct {
	Session *r.Session
}

func (s *StaticStorage) Connect() (err error) {
	s.Session, err = r.Connect(
		r.ConnectOpts{Address: "localhost:28015",
			Database: "test",
			MaxIdle:  10,
			MaxOpen:  20})
	return err
}

func (s *StaticStorage) CreateSchema() error {
	_, err := r.DB("test").TableCreate("champions").RunWrite(s.Session)
	if err != nil {
		return err
	}
	return nil
}

func (s *StaticStorage) AddChampion(champ Champion) (err error) {
	_, err = r.Table("champions").Insert(champ).Run(s.Session)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type VersionList struct {
	Versions []string
}

func StaticVersions() (err error, vl VersionList) {
	path := "https://ddragon.leagueoflegends.com/api/versions.json"
	resp, err := http.Get(path)
	if err != nil {
		return err, vl
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, vl
	}
	json.Unmarshal(body, &vl.Versions)

	return err, vl
}

type LanguageList struct {
	Languages []string
}

func StaticLanguages() (err error, ll LanguageList) {
	path := "https://ddragon.leagueoflegends.com/cdn/languages.json"
	resp, err := http.Get(path)
	if err != nil {
		return err, ll
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ll
	}
	json.Unmarshal(body, &ll.Languages)
	return err, ll
}

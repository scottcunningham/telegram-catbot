package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CatData struct {
	Breeds []string `json:"breeds"`
	ID     string   `json:"id"`
	URL    string   `json:"url"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

var catURL = "https://api.thecatapi.com/v1/images/search"

func getCatURL() (string, error) {
	resp, err := http.Get(catURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	cats := []CatData{}
	err = json.Unmarshal(body, &cats)
	if err != nil {
		return "", err
	}
	if len(cats) != 1 {
		return "", fmt.Errorf("too many cats")
	}
	return cats[0].URL, nil
}

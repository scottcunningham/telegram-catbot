package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CatFact struct {
	Fact string `json:"fact"`
}

var catFactsURL = "https://catfact.ninja/fact"

func randomCatFact() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", catFactsURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fact := CatFact{}
	err = json.Unmarshal(body, &fact)
	if err != nil {
		return "", err
	}
	return fact.Fact, nil
}

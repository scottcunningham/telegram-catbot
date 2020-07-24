package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

type CatFacts struct {
	All []CatFact `json:"all"`
}

type CatFact struct {
	ID      string `json:"_id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	Upvotes int    `json:"upvotes"`
}

var catFactsURL = "https://brianiswu-cat-facts-v1.p.rapidapi.com/facts"
var catFactsKey = os.Getenv("CAT_FACTS_KEY")

func getCatFacts() (*CatFacts, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", catFactsURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-host", "brianiswu-cat-facts-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", catFactsKey)
	req.Header.Add("useQueryString", "true")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	facts := CatFacts{All: []CatFact{}}
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}
	return &facts, nil
}

func (facts *CatFacts) RandomFact() CatFact {
	idx := rand.Intn(len(facts.All))
	return facts.All[idx]
}

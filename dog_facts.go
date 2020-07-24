package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DogFacts struct {
	Facts   []string `json:"facts"`
	Success bool     `json:"success"`
}

var dogFactsURL = "https://dog-api.kinduff.com/api/facts"

func getDogFact() (string, error) {
	resp, err := http.Get(dogFactsURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data DogFacts
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if len(data.Facts) != 1 {
		return "", fmt.Errorf("too many dog facts")
	}
	return data.Facts[0], nil
}

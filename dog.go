package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DogData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var dogURL = "https://dog.ceo/api/breeds/image/random"

func getDogURL() (string, error) {
	resp, err := http.Get(dogURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	dog := DogData{}
	err = json.Unmarshal(body, &dog)
	if err != nil {
		return "", err
	}
	if dog.Status != "success" {
		return "", fmt.Errorf("unsuccessful pup")
	}
	return dog.Message, nil
}

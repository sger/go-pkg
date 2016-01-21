package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlSearchGoAPI = "http://go-search.org/api"

//SearchResponse contains an array of ResultsData
type SearchResponse struct {
	Results []ResultsData `json:"hits"`
}

//ResultsData contains information about the Go package
type ResultsData struct {
	Name       string `json:"name"`
	Package    string `json:"package"`
	Author     string `json:"author"`
	ProjectURL string `json:"projecturl"`
}

func search(query string) (*SearchResponse, error) {
	url := fmt.Sprintf(urlSearchGoAPI+"?action=search&q=%s", query)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `application/json`)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result SearchResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

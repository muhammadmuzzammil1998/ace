package main

type response struct {
	Version string       `json:"version"`
	Data    responseData `json:"data"`
}

type responseData struct {
	URL         string `json:"url"`
	OldHash     string `json:"oldHash"`
	NewHash     string `json:"newHash"`
	Source      string `json:"source"`
	LastUpdated string `json:"lastUpdated"`
	IsForced    bool   `json:"isForced"`
}

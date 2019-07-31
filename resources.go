package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Resources - base structure
type resources struct {
	Version string  `json:"version"`
	LastRun lastRun `json:"lastRun"`
	Data    []data  `json:"data"`
}

// Data - individual records
type data struct {
	URL         string `json:"url"`
	Hash        string `json:"hash"`
	Action      string `json:"action"`
	Location    string `json:"location"`
	LastUpdated string `json:"lastUpdated"`
	IsForced    bool   `json:"isForced"`
}

// LastRun - information about last run
type lastRun struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (rs *resources) add(r data) *resources {
	rs.Data = append(rs.Data, r)
	return rs
}

func (rs *resources) remove(u string) bool {
	temp := &resources{
		LastRun: rs.LastRun,
		Version: rs.Version,
	}
	var found bool
	for _, data := range rs.Data {
		if data.URL == u {
			found = true
			continue
		}
		temp.add(data)
	}
	*rs = *temp
	return found
}

func (rs *resources) read() *resources {
	data, err := ioutil.ReadFile(flags.datafile)
	check(err)

	r := &resources{}
	err = json.Unmarshal(data, r)
	check(err)
	return r
}

func (rs *resources) write() *resources {
	rs.Version = version
	if strings.Trim(rs.LastRun.Start, " ") == "" {
		rs.LastRun.Start = "never"
		rs.LastRun.End = "never"
	}
	j, err := json.MarshalIndent(rs, "", "  ")
	check(err)
	file, err := os.OpenFile(flags.datafile, os.O_CREATE|os.O_WRONLY, 0600)
	check(err)
	defer file.Close()

	file.WriteString(string(j))
	return rs
}

func (rs *resources) overwrite() {
	err := os.Remove(flags.datafile)
	check(err)
	rs.write()
}

func (d *data) getSource() (string, error) {
	resp, err := http.Get(d.URL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("Page is unreachable. Status: " + resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

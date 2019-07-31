package main

type flagVars struct {
	url      string
	location string
	action   string
	datafile string
	add      bool
	forceAdd bool
	remove   bool
	version  bool
	interval float64
}

var (
	version string
	home    string
	flags   flagVars
)

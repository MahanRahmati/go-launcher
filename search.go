package main

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func fuzzySearch(query string, applications []Application) []Application {
	apps := []Application{}
	for _, app := range applications {
		if fuzzy.MatchFold(query, app.Name) {
			apps = append(apps, app)
		}
	}
	return apps
}

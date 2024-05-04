package main

import (
	"log"
)

func main() {
	os := getOS()
	applications, err := getApplications(os)
	if err != nil {
		log.Fatal(err)
	}
	query := getUserInput()
	filteredApplications := fuzzySearch(query, applications)
	first_app := filteredApplications[0]
	out, errout, err := openApplication(os, first_app)
	if err != nil {
		log.Fatal(err)
	}
	if out != "" {
		log.Print("stdout: " + out)
	}
	if errout != "" {
		log.Print("stderr: " + errout)
	}
}

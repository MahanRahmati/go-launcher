package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type Application struct {
	Name string
	path string
	cmd  string
}

func getApplications(os string) ([]Application, error) {
	switch os {
	case Darwin:
		applications, err := getMacApplications()
		if err != nil {
			return nil, err
		}
		return applications, nil
	case Linux:
		applications, err := getLinuxApplications()
		if err != nil {
			return nil, err
		}
		return applications, nil
	case Windows:
		return nil, fmt.Errorf("windows is not supported")
	default:
		return nil, fmt.Errorf("operating system not supported")
	}
}

func getMacApplications() ([]Application, error) {
	applications := []Application{}
	locations := []string{
		"/System/Applications/",
		"/Applications/",
	}
	for _, location := range locations {
		files, err := readDirectory(location)
		if err != nil {
			return nil, err
		}
		apps := getMacApps(files, location)
		applications = append(applications, apps...)
	}
	return applications, nil
}

func readDirectory(path string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getMacApps(files []fs.DirEntry, applicationPath string) []Application {
	applications := []Application{}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".app") {
			applications = append(
				applications,
				Application{
					Name: filename,
					path: applicationPath,
				},
			)
		} else if file.IsDir() {
			newPath := applicationPath + filename + "/"
			subfiles, err := os.ReadDir(newPath)
			if err != nil {
				continue
			}
			apps := getMacApps(subfiles, newPath)
			applications = append(applications, apps...)
		}
	}
	return applications
}

func getLinuxApplications() ([]Application, error) {
	applications := []Application{}
	locations := []string{
		"/usr/local/share/applications/",
		"/usr/share/applications/",
		"~/.local/share/applications",
	}
	for _, location := range locations {
		files, err := readDirectory(location)
		if err != nil {
			return nil, err
		}
		apps := getLinuxApps(files, location)
		applications = append(applications, apps...)
	}
	return applications, nil
}

func getLinuxApps(files []fs.DirEntry, applicationPath string) []Application {
	applications := []Application{}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".desktop") {
			data, err := os.ReadFile(applicationPath + filename)
			if err != nil {
				log.Fatal(err)
			}
			str := string(data)
			if strings.Contains(str, "Exec=") {
				start := strings.Index(str, "Exec=") + 5
				end := strings.Index(str[start:], "\n")
				if end == -1 {
					end = len(str)
				}
				end += start
				cmd := str[start:end]
				applications = append(
					applications,
					Application{
						Name: filename,
						cmd:  cmd,
					},
				)
			}

		}
	}
	return applications
}

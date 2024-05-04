package main

import (
	"runtime"
	"strings"
)

const (
	Linux   string = "linux"
	Darwin  string = "darwin"
	Windows string = "windows"
)

func getOS() string {
	os := runtime.GOOS
	if strings.Contains(os, "darwin") {
		return Darwin
	} else if strings.Contains(os, "linux") {
		return Linux
	} else if strings.Contains(os, "windows") {
		return Windows
	} else {
		return ""
	}
}

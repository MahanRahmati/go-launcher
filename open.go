package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func openApplication(os string, application Application) (out string, errout string, error error) {
	switch os {
	case Darwin:
		return runMacApplication(application)
	case Linux:
		return runLinuxApplication(application)
	case Windows:
		return "", "", fmt.Errorf("windows is not supported")
	default:
		return "", "", fmt.Errorf("operating system not supported")
	}
}

func Shellout(command string) (out string, errout string, error error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func runMacApplication(application Application) (out string, errout string, error error) {
	out, errout, err := Shellout("open \"" + application.path + application.Name + "\"")
	return out, errout, err
}

func runLinuxApplication(application Application) (out string, errout string, error error) {
	out, errout, err := Shellout(application.cmd)
	return out, errout, err
}

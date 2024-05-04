package main

import (
	"fmt"
	"os"
)

func getUserInput() string {
	arg := os.Args[1]
	if arg != "" {
		return arg
	}
	var input string
	fmt.Scan(&input)
	return input
}

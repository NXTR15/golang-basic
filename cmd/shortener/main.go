package main

import (
	"fmt"
	"os"
)

// Package only Variable (private variable)
var appName = "url-shortener"

// Version variable (public variable)
const Version = "0.1.0"

func main() {
	var argument = os.Args[1:]

	if len(argument) < 1 {
		fmt.Fprintf(os.Stderr, "butuh 1 argumen (url)\n")
		os.Exit(1)
	}

	url := argument[0]

	fmt.Printf("%s %s\n", appName, Version)
	fmt.Printf("url=%s len=%d\n", url, len(url))
}

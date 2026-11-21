package main

import "fmt"

// Package only Variable (private variable)
var appName = "url-shortener"

// Version variable (public variable)
const Version = "0.1.0"

func main() {
	var links int       // Golang has no null value, so this variable value is 0 (integer)
	var lastCode string // Default value is empty string ("")
	ready := true       // Short version of (=) and also declare new variable while fill it in

	//unused := 42 // This variable will cause error if not used, so unused variable need to be deleted

	fmt.Printf("%s version %s\n", appName, Version) // Same like printf in java where replace %s using variable
	fmt.Printf("links=%d lastCode=%q ready=%t\n", links, lastCode, ready)
}

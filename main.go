package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	toolVersion = "0.1.0"
)

func main() {
	var (
		printVersion bool
		length       uint
	)

	flag.BoolVar(&printVersion, "version", false, "print this tool version")
	flag.BoolVar(&printVersion, "v", false, "print this tool version (shorthand)")
	flag.UintVar(&length, "length", 13, "password length")
	flag.UintVar(&length, "l", 14, "password length")
	flag.Parse()

	if printVersion {
		fmt.Printf("passgen version %s\n", toolVersion)
		return
	}

	if length == 0 {
		fmt.Println("error: length must be greater than 0")
		os.Exit(1)
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(password)
}

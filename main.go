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
		length       int
	)

	flag.BoolVar(&printVersion, "version", false, "print passgen version")
	flag.BoolVar(&printVersion, "v", false, "print passgen version (shorthand)")
	flag.IntVar(&length, "length", 16, "password length")
	flag.IntVar(&length, "l", 16, "password length")
	flag.Parse()

	if printVersion {
		fmt.Printf("passgen version %s\n", toolVersion)
		os.Exit(0)
	}

	if length <= 0 {
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

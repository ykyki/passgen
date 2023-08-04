package main

import (
	"flag"
	"fmt"
)

var (
	toolVersion = "0.1.0"
)

func main() {
	var (
		printVersion bool
	)

	flag.BoolVar(&printVersion, "version", false, "print this tool version")
	flag.BoolVar(&printVersion, "v", false, "print this tool version (shorthand)")
	flag.Parse()

	if printVersion {
		fmt.Printf("passgen version %s\n", toolVersion)
		return
	}

	fmt.Println("passgen (TODO)")
}

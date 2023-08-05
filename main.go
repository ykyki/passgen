package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	toolVersion = "0.1.0"
)

func main() {
	cli := &cli{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.run(os.Args))
}

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) run(args []string) int {
	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.SetOutput(c.errStream)

	var (
		printVersion bool
		length       int
	)

	f.BoolVar(&printVersion, "version", false, "print passgen version")
	f.BoolVar(&printVersion, "v", false, "print passgen version (shorthand)")
	f.IntVar(&length, "length", 16, "password length")
	f.IntVar(&length, "l", 16, "password length")

	err := f.Parse(args[1:])
	if err != nil {
		// fmt.Fprintf(c.errStream, "failed to parse arguments\n%s\n", err)
		fmt.Fprintln(c.errStream, "failed to parse arguments")
		return 1
	}

	if printVersion {
		fmt.Fprintf(c.outStream, "passgen version %s\n", toolVersion)
		return 0
	}

	if length <= 0 {
		fmt.Fprintf(c.errStream, "length must be greater than 0\n")
		return 0
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Fprintf(c.errStream, "failed to generate password\n%s\n", err)
		return 1
	}

	fmt.Fprintf(c.outStream, "%s\n", password)

	return 0
}

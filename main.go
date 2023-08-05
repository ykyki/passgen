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
		printVersion                                               bool
		length                                                     int
		capitalLetterFlag, smallLetterFlag, numberFlag, symbolFlag bool
	)

	f.BoolVar(&printVersion, "version", false, "print passgen version")
	f.BoolVar(&printVersion, "v", false, "print passgen version (shorthand)")
	f.IntVar(&length, "length", 16, "password length")
	f.IntVar(&length, "l", 16, "password length")
	f.BoolVar(&capitalLetterFlag, "A", false, "include capital letters")
	f.BoolVar(&smallLetterFlag, "a", false, "include small letters")
	f.BoolVar(&numberFlag, "n", false, "include numbers")
	f.BoolVar(&symbolFlag, "s", false, "include symbols")

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

	password, err := generatePassword(
		length,
		decideLetterOption(capitalLetterFlag, smallLetterFlag, numberFlag, symbolFlag),
	)
	if err != nil {
		fmt.Fprintf(c.errStream, "failed to generate password\n%s\n", err)
		return 1
	}

	fmt.Fprintf(c.outStream, "%s\n", password)

	return 0
}

type letterOption struct {
	capitalLetter, smallLetter, number, symbol bool
}

func decideLetterOption(capitalLetterFlag, smallLetterFlag, numberFlag, symbolFlag bool) letterOption {
	if capitalLetterFlag || smallLetterFlag || numberFlag || symbolFlag {
		return letterOption{
			capitalLetter: capitalLetterFlag,
			smallLetter:   smallLetterFlag,
			number:        numberFlag,
			symbol:        symbolFlag,
		}
	}

	return letterOption{
		capitalLetter: true,
		smallLetter:   true,
		number:        true,
		symbol:        true,
	}
}

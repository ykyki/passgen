package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	passgenVersion = "0.2.0"
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

	f.BoolVarP(&printVersion, "version", "v", false, "print passgen version")
	f.IntVarP(&length, "length", "l", 32, "password length")
	f.BoolVarP(&capitalLetterFlag, "capital", "A", false, "include capital letters")
	f.BoolVarP(&smallLetterFlag, "small", "a", false, "include small letters")
	f.BoolVarP(&numberFlag, "number", "n", false, "include numbers")
	f.BoolVarP(&symbolFlag, "symbol", "s", false, "include symbols")
	f.SortFlags = false

	err := f.Parse(args[1:])
	if err != nil {
		// fmt.Fprintf(c.errStream, "failed to parse arguments\n%s\n", err)
		fmt.Fprintln(c.errStream, "failed to parse arguments")
		return 1
	}

	if printVersion {
		fmt.Fprintf(c.outStream, "passgen version %s\n", passgenVersion)
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

package main

import (
	"flag"
	"log"
	"os"
)

var (
	damerauLevenshteinDistance = flag.Int("dist", 20, "Minimal Damerau-Levenshtein distance to match")
	top                        = flag.Int("top", 10, "How many lines will be shown")
	debug                      = flag.Bool("debug", false, "Debug output. WARNING! this cause huge output")
)

func main() {
	flag.Parse()

	input, err := detectInput()
	if err != nil {
		log.Fatalf("Please use stdin or first parameter as input.\nErr: '%s'\n", err)
	}
	it := make(results, 0)
	it.processInput(input, *debug)

	err = it.print(os.Stdout, *top, consoleOut)
	if err != nil {
		log.Fatalf("Print error.\nErr: '%s'\n", err)
	}
}

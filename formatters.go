package main

import (
	"fmt"
)

func consoleOut(r result) []byte {
	delimeter := "----------------------------------\n"
	return []byte(fmt.Sprintf("%s\t- Line Example: '%s'\n\tMatches: %d\n\tAverage Levenstein distance: %f\n", delimeter, r.source, r.matches, r.avgLevDist))
}

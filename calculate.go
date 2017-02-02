package main

import (
	"bufio"
	"io"
	"log"

	"github.com/masatana/go-textdistance"
)

func (r *results) processInput(input io.Reader, debug bool) {
	var (
		match          bool
		bestMatchIndx  int
		bestMatchCoeff int
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		match = false
		bestMatchIndx = 0
		bestMatchCoeff = *damerauLevenshteinDistance
		currString := scanner.Text()
		for i, res := range *r {
			dld := textdistance.DamerauLevenshteinDistance(res.source, currString)
			if dld > *damerauLevenshteinDistance {
				continue
			}
			match = true
			if dld < bestMatchCoeff {
				bestMatchCoeff = dld
				bestMatchIndx = i
				if debug {
					log.Printf(" Log line '%s' seems like '%s' because Damerau-Levenshtein distance is '%d' \n", currString, res.source, dld)
				}
			}
		}
		if !match {
			*r = append(*r, result{source: currString})
			continue
		}
		if (*r)[bestMatchIndx].matches != 0 {
			(*r)[bestMatchIndx].avgLevDist = ((*r)[bestMatchIndx].avgLevDist + float32(bestMatchCoeff)) / 2
		} else {
			(*r)[bestMatchIndx].avgLevDist = float32(bestMatchCoeff)
		}
		(*r)[bestMatchIndx].matches++
	}
}

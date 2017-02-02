package main

import (
	"io"
	"sort"

	"github.com/pkg/errors"
)

type results []result

type result struct {
	source     string
	matches    int64
	avgLevDist float32
}

func (r *results) Len() int           { return len(*r) }
func (r *results) Less(i, j int) bool { return (*r)[i].matches > (*r)[j].matches }
func (r *results) Swap(i, j int)      { (*r)[i], (*r)[j] = (*r)[j], (*r)[i] }

func (r *results) print(out io.Writer, top int, fmtr func(r result) []byte) error {
	sort.Sort(r)
	for i := 0; i < top; i++ {
		if i+1 > r.Len() {
			break
		}
		_, err := out.Write(fmtr((*r)[i]))
		if err != nil {
			return errors.Wrap(err, "Coudn't print result")
		}
	}
	return nil
}

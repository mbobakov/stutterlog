package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_results_processInput(t *testing.T) {
	tests := []struct {
		name   string
		expect results
		input  io.Reader
	}{
		{name: "testCase1", expect: results{result{source: "test", matches: 3, avgLevDist: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open("testdata/" + tt.name)
			if err != nil {
				t.Error(err)
			}
			tt.input = f
			got := new(results)
			got.processInput(tt.input, false)
			if !reflect.DeepEqual(*got, tt.expect) {
				t.Errorf("%s:\nGot '%+v'\nExpect: '%+v'", tt.name, *got, tt.expect)
			}
		})
	}
}

package main

import (
	"io"
	"os"

	"flag"

	"github.com/pkg/errors"
)

// detectInput returns io.Reader with input data
// TODO(mbobakov): Close file gracefully
func detectInput() (io.Reader, error) {
	stats, err := os.Stdin.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "file.Stat() in detectInput")
	}

	if stats.Mode()&os.ModeNamedPipe != 0 {
		return os.Stdin, nil
	}
	if len(flag.Args()) < 1 {
		return nil, errors.New("No inputs found")
	}
	if len(flag.Args()) > 1 {
		return nil, errors.Errorf("Multiple input files is not supported. Your choise is '%v'", flag.Args()[1:])
	}
	f, err := os.Open(flag.Args()[0])
	if err != nil {
		return nil, errors.Wrap(err, "Coudn't open file")
	}
	return f, nil
}

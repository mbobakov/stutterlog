
[![Go Report Card](https://goreportcard.com/badge/github.com/mbobakov/stutterlog)](https://goreportcard.com/report/github.com/mbobakov/stutterlog) [![Build Status](https://travis-ci.org/mbobakov/stutterlog.svg?branch=master)](https://travis-ci.org/mbobakov/stutterlog)

# stutterlog
'''stutterlog''' is dead simple tool for search for dublicates in the log.
It's use [Damerau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau–Levenshtein_distance) for finding almost same log lines.

### Install
``` go get https://github.com/mbobakov/stutterlog```

or download via releases page
### Usage
    Usage of ./stutterlog:
      -debug
            Debug output. WARNING! this cause huge output
      -dist int
            Minimal Damerau-Levenshtein distance to match (default 20)
      -top int
            How many lines will be shown (default 10)
### Examples
Let's me show how it works. For example we have log file like this:

    2015-03-26T01:27:38-04:00 debug trash message with tr@sH
    2015-03-26T01:27:39-04:00 info status of id=1 is OK
    2015-03-26T01:27:42-04:00 info status of id=2 is OK
    2015-03-26T01:27:44-04:00 debug trash message with t4H
    2015-03-26T01:27:44-04:00 info status of id=3 is OK
    2015-03-26T01:27:45-04:00 debug trash message with t7H

You can pass logs via stdin or as argument:

    mbobakov$> cat test | ./stutterlog
    ----------------------------------
            - Line Example: '2015-03-26T01:27:38-04:00 debug trash message with tr@sH'
            Matches: 2
            Average Levenstein distance: 5.000000
    ----------------------------------
            - Line Example: '2015-03-26T01:27:39-04:00 info status of id=1 is OK'
            Matches: 2
            Average Levenstein distance: 3.000000

### Performance
On 2 GHz Intel Core i7 / 8 GB 1600 MHz DDR3 / SSD. 1000 log lines with 512 symbols for each

    mbobakov$> time ./stutterlog 1000linedFile.log > /dev/null
            1.07 real         1.10 user         0.02 sys

### Contibuting
Go >= 1.7 is requirement. This tool use [dep](https://github.com/golang/dep) for dependency management. Pull Request are welcome. ;) 

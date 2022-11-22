package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Initialize the command-line arguments
	var (
		prefix    string
		delimiter string
		verbose   bool
	)

	var REGIONAL_INDICATOR = "regional_indicator_"

	flag.StringVar(&prefix, "p", REGIONAL_INDICATOR, "Prefix to be attached to message.")
	// NOTE: delimiter's default value is ZWSP(Zero Witch Space).
	flag.StringVar(&delimiter, "d", "​", "Delimiter to concatenate the text")
	flag.BoolVar(&verbose, "v", false, "Verbose mode?")

	flag.Parse()

	// Initialize other things
	var validPattern = regexp.MustCompile(`[a-zA-Z]`)
	var matchPattern func(string) bool

	if prefix == REGIONAL_INDICATOR {
		matchPattern = func(t string) bool {
			return validPattern.MatchString(t)
		}
	} else {
		matchPattern = func(t string) bool {
			return true
		}
	}

	// GO GO GO
	var args = strings.Join(flag.Args(), " ")
	var text = strings.Split(args, "")

	if verbose {
		fmt.Printf("Parameter prefix   : \"%s\"\n", prefix)
		fmt.Printf("Parameter delimtier: \"%s\"\n", delimiter)
	}

	var texts []string

	for _, t := range text {
		if t != " " && matchPattern(t) {
			texts = append(texts, ":"+prefix+strings.ToLower(t)+":")
		} else {
			texts = append(texts, t)
		}
	}

	fmt.Printf("%s\n", strings.Join(texts, delimiter))
}

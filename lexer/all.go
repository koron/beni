package lexer

import (
	"log"
	"regexp"
)

// All factories.
var All = []Factory{
	Go,
	Java,
}

func matchFilename(t, p string) bool {
	matched, err := regexp.MatchString(p, t)
	if err != nil {
		log.Printf("lexer has illegal pattern: %s", p)
		return false
	}
	return matched
}

func matchByFilenames(s string, n Info) bool {
	for _, p := range n.Filenames {
		if matchFilename(s, p) {
			return true
		}
	}
	for _, p := range n.AliasFilenames {
		if matchFilename(s, p) {
			return true
		}
	}
	return false
}

// Find returns matched Theme.
func Find(s string) Factory {
	for _, f := range All {
		if matchByFilenames(s, f.Info()) {
			return f
		}
	}
	return nil
}

// Package regexplus provides additional convenience functions over regexp.
package regexplus

import (
	"regexp"
)

// A Map represents the key-value pairs generated by a regexp
// named group string matching. The keys are the names of named
// groups and the values are the matched strings.
type Map map[string]string

// FindNamedSubmatches uses re to find named submatches of s, and returns
// a Map from keys (names of groups) to values (matches).
//
// Example:
//   re := regexp.MustCompile(`(?P<foo>[a-z]+)-(?P<bar>[0-9]+)`)
//   m := FindNamedSubmatches(re, "abc-123")
//   m == Map{"foo":"abc", "bar":"123"}
func FindNamedSubmatches(re *regexp.Regexp, s string) Map {
	names := re.SubexpNames()
	groups := re.FindStringSubmatch(s)

	if len(groups) == 0 {
		return nil
	}

	m := make(Map, len(names)-1)
	for i := 1; i < len(names); i++ {
		m[names[i]] = groups[i]
	}
	return m
}

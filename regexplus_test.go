package regexplus

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Map(t *testing.T) {
	try := func(r *regexp.Regexp, input string, exp Map) {
		m := FindNamedSubmatches(r, input)
		if len(exp) == 0 {
			require.Equal(t, 0, len(m))
		} else {
			require.Equal(t, exp, m)
		}
	}

	re0 := regexp.MustCompile(`fooBar`)
	try(re0, "fooBar", nil)
	try(re0, "noMatch", nil)

	re1 := regexp.MustCompile(`[a-z]+-(?P<nums>[0-9]+)`)
	try(re1, "abc-123", Map{
		"nums": "123",
	})
	try(re1, "noMatch", nil)

	re2 := regexp.MustCompile(`(?P<a>[\d]+)-(?P<b>[\d]+)-(?P<c>[\d])`)
	try(re2, "12-34-56", Map{
		"a": "12",
		"b": "34",
		"c": "5",
	})

	re4 := regexp.MustCompile(`(?P<a>[\d]*)-(?P<b>[\d]*)-(?P<c>[\d])`)
	try(re4, "12--56", Map{
		"a": "12",
		"b": "",
		"c": "5",
	})
}

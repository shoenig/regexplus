package regexplus

import (
	"reflect"
	"regexp"
	"testing"
)

func eqMap(t *testing.T, exp, got Map) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func eqInt(t *testing.T, exp, got int) {
	if exp != got {
		t.Fatalf("expected %d, got %d", exp, got)
	}
}

func Test_Map(t *testing.T) {
	try := func(r *regexp.Regexp, input string, exp Map) {
		m := FindNamedSubmatches(r, input)
		if len(exp) == 0 {
			eqInt(t, 0, len(m))
		} else {
			eqMap(t, exp, m)
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

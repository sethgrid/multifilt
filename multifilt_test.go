package multifilt

import (
	"bytes"
	"testing"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		in          []byte
		filter      []byte
		requireFull bool
		expected    []byte
	}{
		{
			in:          []byte("abc\nbcd\ncde"),
			filter:      []byte("bc\n"),
			requireFull: false,
			expected:    []byte("cde\n"),
		},
		{
			in:          []byte("abba\ndabba\ndo\nraz\nI pitty the fool\n"),
			filter:      []byte("ba\nfoo\n"),
			requireFull: false,
			expected:    []byte("do\nraz\n"),
		},
		{
			in:          []byte("abba\ndabba\ndo\nraz\nI pitty the fool\n"),
			filter:      []byte("do\nraz\nabba\nba\n"),
			requireFull: true,
			expected:    []byte("dabba\nI pitty the fool\n"),
		},
	}
	for _, test := range cases {
		in := bytes.NewBuffer(test.in)
		filter := bytes.NewBuffer(test.filter)

		var b []byte
		out := bytes.NewBuffer(b)

		err := Filter(in, filter, out, test.requireFull)
		if err != nil {
			t.Errorf("got an error - %v", err)
		}

		if bytes.Compare(out.Bytes(), test.expected) != 0 {
			t.Errorf("got:\n%s\n\nwant:\n%s", out.Bytes(), test.expected)
		}
	}
}

func TestIsMatch(t *testing.T) {
	cases := []struct {
		haystack    []byte
		needle      []byte
		requireFull bool
		shouldMatch bool
	}{
		{
			haystack:    []byte("abcdef"),
			needle:      []byte("cde"),
			requireFull: false,
			shouldMatch: true,
		},
		{
			haystack:    []byte("abcdef"),
			needle:      []byte("cde"),
			requireFull: true,
			shouldMatch: false,
		},
	}

	for _, test := range cases {
		matched, err := isMatch(test.haystack, test.needle, test.requireFull)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if matched != test.shouldMatch {
			t.Errorf("%s <=> %s, got match? %t, want match? %t", test.haystack, test.needle, matched, test.shouldMatch)
		}
	}
}

package main

import "testing"

func TestSize(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{-1, "negative"},
		{0, "zero"},
		{5, "small"},
		{50, "big"},
		{500, "huge"},
		{10000, "enormous"},
	}
	for _, test := range tests {
		s := size(test.in)
		if s != test.out {
			t.Errorf("Size(%d)=%s; expected %s", test.in, s, test.out)
		}
	}
}

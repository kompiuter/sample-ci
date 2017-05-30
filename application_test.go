package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("num", "5")
	req.URL.RawQuery = q.Encode()

	rw := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rw, req)

	want := "5 is small"
	got := rw.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %q, want %q", got, want)
	}
}

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
			t.Errorf("Size(%d)=%s; want %s", test.in, s, test.out)
		}
	}
}

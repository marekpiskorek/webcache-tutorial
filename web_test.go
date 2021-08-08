package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/webcache/webcache"
)

func TestEmptyDatabase(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/webcache", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	webcache.CachedWebpageHandler(res, req)

	exp := "No cache found"
	act := res.Body.String()

	if exp != act {
		t.Fatalf("Expected >>%s<< got >>%s<<", exp, act)
	}
}

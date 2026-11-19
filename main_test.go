package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// /version deve restituire qualcosa
func TestVersion(t *testing.T) {
	rr := httptest.NewRecorder()
	Version(rr, httptest.NewRequest("GET", "/version", nil))
	if strings.TrimSpace(rr.Body.String()) == "" {
		t.Fatal("/version ha risposto vuoto")
	}
}

// ci devono essere citazioni, ognuna con testo e autore
func TestQuotes(t *testing.T) {
	if len(quotes) == 0 {
		t.Fatal("nessuna citazione")
	}
	for _, q := range quotes {
		if q.Quotation == "" || q.Author == "" {
			t.Errorf("citazione %d incompleta", q.Id)
		}
	}
}

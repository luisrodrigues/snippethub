package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"snippethub.luisrodrigues.io/internal/assert"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)
	rs := rr.Result()
	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}

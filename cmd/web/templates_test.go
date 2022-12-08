package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	expected := "17 Mar 2022 at 10:15"
	tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != expected {
		t.Errorf("got %q, expected %q", hd, expected)
	}
}

package main

import "testing"

func TestNtpServerIsCorrect(t *testing.T) {
	ntpServer = "Hello, Gopher!"
	if err := getTime(); err == nil {
		t.Error(err)
	}
}
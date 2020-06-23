package main

import (
	"io/ioutil"
	"testing"
)

func TestGetResume(test *testing.T) {
	// Check the response body is what we expect.
	data, _ := ioutil.ReadFile("./expected.json")
	expected := string(data)
	resume := getResume()
	actual := resume.formatResume()
	if expected != actual {
		test.Errorf("The resume did not match expected result: got %v want %v", actual, expected)
	}
}

package main

import (
	"strings"
	"testing"
)

func TestFormatLog_JSON(t *testing.T) {
	input := `{"key":"value","nested":{"num":123}}`
	expected := `{
  "key": "value",
  "nested": {
    "num": 123
  }
}`
	output := formatLog(input)
	if strings.TrimSpace(output) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}

func TestFormatLog_NonJSON(t *testing.T) {
	input := "This is a plain log entry"
	output := formatLog(input)
	if output != input {
		t.Errorf("Expected:\n%s\nGot:\n%s", input, output)
	}
}

package parser

import "testing"

func TestParseLevel(t *testing.T) {
	line := "2026-04-23 INFO Server started"

	level := ParseLevel(line)

	if level != "INFO" {
		t.Errorf("expected INFO, got %s", level)
	}
}

func TestParseErrorMessage(t *testing.T) {
	line := "2026-04-23 ERROR Database failed"

	msg := ParseErrorMessage(line)

	expected := "Database failed"

	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

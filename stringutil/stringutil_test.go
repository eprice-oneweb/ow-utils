package stringutil

import "testing"

func TestMatchScidPatternValidScid(t *testing.T) {
	match, err := MatchScidPattern("SL1234")
	if err != nil {
		t.Error("Error should be nil")
	}
	if !match {
		t.Error("SL1234 should match the pattern")
	}
}

func TestMatchScidPatternInvalidScid(t *testing.T) {
	match, err := MatchScidPattern("SL123")
	if err != nil {
		t.Error("Error should be nil")
	}
	if match {
		t.Error("SL123 should not match the pattern")
	}
}

func TestContainsDoesContain(t *testing.T) {
	s := []string{"a", "b", "c"}
	if !Contains(s, "b") {
		t.Error("s should contain b")
	}
}

func TestContainsDoesNotContain(t *testing.T) {
	s := []string{"a", "b", "c"}
	if Contains(s, "d") {
		t.Error("s should not contain d")
	}
}

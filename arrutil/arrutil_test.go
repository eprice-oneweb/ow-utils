package arrutil

import "testing"

func TestContainsDoesContain(t *testing.T) {
	var s []interface{}
	s = append(s, "a", "b", "c")
	if !Contains(s, "b") {
		t.Error("s should contain b")
	}
}

func TestContainsDoesNotContain(t *testing.T) {
	var s []interface{}
	s = append(s, "a", "b", "c")
	if Contains(s, "d") {
		t.Error("s should not contain d")
	}
}

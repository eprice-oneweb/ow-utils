package logger

import (
	"testing"
)

func TestGenerateLogger(t *testing.T) {
	Logger = generateLogger()
	if Logger == nil {
		t.Error("Logger should not be nil")
	}
}

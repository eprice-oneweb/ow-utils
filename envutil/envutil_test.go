package envutil

import (
	"os"
	"testing"
)

func TestGetFloat64EnvOrDefault(t *testing.T) {
	os.Setenv("TEST_FLOAT64", "3.14")
	if GetFloat64EnvOrDefault("TEST_FLOAT64", 0) != 3.14 {
		t.Error("GetFloat64EnvOrDefault() should return 3.14")
	}
}

func TestGetInt64EnvOrDefault(t *testing.T) {
	os.Setenv("TEST_INT64", "314")
	if GetInt64EnvOrDefault("TEST_INT64", 0) != 314 {
		t.Error("GetInt64EnvOrDefault() should return 314")
	}
}

func TestGetIntEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_INT", "314")
	if GetIntEnvOrDefault("TEST_INT", 0) != 314 {
		t.Error("GetIntEnvOrDefault() should return 314")
	}
}	

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_ENV", "314")
	if GetEnvOrDefault("TEST_ENV", "0") != "314" {
		t.Error("GetEnvOrDefault() should return 314")
	}
}	

func TestGetBoolEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")
	if GetBoolEnvOrDefault("TEST_BOOL", false) != true {
		t.Error("GetBoolEnvOrDefault() should return true")
	}
}

func TestGetBoolEnvOrDefault2(t *testing.T) {
	os.Setenv("TEST_BOOL", "false")
	if GetBoolEnvOrDefault("TEST_BOOL", true) != false {
		t.Error("GetBoolEnvOrDefault() should return false")
	}
}
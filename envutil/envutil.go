package envutil

import (
	"os"
	"strconv"
)
func GetPathEnvWithSlash(env string) string {
	path := os.Getenv(env)
	if path == "" {
		return ""
	} else if path[len(path)-1:] != "/" {
		return path + "/"
	}
	return path
}

func GetFloat64EnvOrDefault(variable string, defaultValue float64) float64 {
	floatenv, err := strconv.ParseFloat(os.Getenv(variable), 64)
	if err != nil {
		return defaultValue
	}
	return floatenv
}

func GetInt64EnvOrDefault(variable string, defaultValue int64) int64 {
	intenv, err := strconv.ParseInt(os.Getenv(variable), 10, 32)
	if err != nil {
		return defaultValue
	}
	return intenv
}

func GetIntEnvOrDefault(variable string, defaultValue int) int {
	intenv, err := strconv.Atoi(os.Getenv(variable))
	if err != nil {
		return defaultValue
	}
	return intenv
}

func GetEnvOrDefault(variable, defaultValue string) string {
	env := os.Getenv(variable)
	if env == "" {
		return defaultValue
	}
	return env
}

func GetBoolEnvOrDefault(variable string, defaultValue bool) bool {
	boolenv, err := strconv.ParseBool(os.Getenv(variable))
	if err != nil {
		return defaultValue
	}
	return boolenv
}

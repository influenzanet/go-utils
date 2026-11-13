package configs

import (
	"os"
	"strconv"
)

// RequireEnv gets value for an environment variable, and raises an error if variable is not defined or if value is empty if `allowEmpty` is false.
// message is sent using internal error logger, configurable using `SetErrorLogger`
func RequireEnv(name string, allowEmpty bool) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		Fatalf("Envirnoment variable '%s' must be defined", name)
	}
	if value == "" {
		if !allowEmpty {
			Fatalf("Envirnoment variable '%s' must have a non empty value", name)
		}
	}
	return value
}

// RequireEnvValue gets a value from env and raises a message if variable is not defined or value is empty.
// This is to be used to require a non empty value for an environment variable.
// message is sent using internal error logger, configurable using `SetErrorLogger`
func RequireEnvValue(name string) string {
	return RequireEnv(name, false)
}

// RequireEnvSet gets a value from env and raises a message if the variable is not defined. It allows an empty value.
// message is sent using internal error logger, configurable using `SetErrorLogger`
func RequireEnvSet(name string) string {
	return RequireEnv(name, true)
}

// Get Env variable with a default value
func GetEnvOr(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

// GetEnvInt get integer value from env name, if fails raise Fatal logger, if empty use default value
func GetEnvInt(name string, def int) int {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		Fatalf(" error in `%s`, expect integer value: %s", name, err)
	}
	return i
}

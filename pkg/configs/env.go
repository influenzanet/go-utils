package configs

import (
	"os"
	"strconv"
)

// RequireEnv get a value from env and raise a message if value is empty
// message is sent using internal error logger, configurable using `SetErrorLogger`
func RequireEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		Fatalf("Expect a value for env '%s'", name)
	}
	return value
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

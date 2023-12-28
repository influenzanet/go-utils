package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// ParseDuration parse value as a time duration format or if no unit is provided use the default provided one
func ParseDuration(value string, defaultUnit string) (time.Duration, error) {
	// Provided value is only a numeric string
	if v, err := strconv.Atoi(value); err == nil {
		value = fmt.Sprintf("%d"+defaultUnit, v)
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		return time.Duration(0), fmt.Errorf("invalid time duration '%s' : %s", value, err.Error())
	}
	return d, nil
}

// ParseEnvDuration gets value from environment variable (name), and parse it as duration value or with default unit
// if only numeric value is provided.
func ParseEnvDuration(name string, defaultValue time.Duration, defaultUnit string) time.Duration {
	value := os.Getenv(name)
	if value == "" {
		Infof("%s : not provided using default value %s", name, defaultValue)
		return defaultValue
	}
	d, err := ParseDuration(value, defaultUnit)
	if err != nil {
		Errorf("%s : unexpected error - default value used, %s", name, err.Error())
		return defaultValue
	}
	Infof("%s : using value %s", name, d)
	return d
}

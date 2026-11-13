package configs

import (
	"fmt"
	"os"
	"strconv"
)

type DBConfig struct {
	URI             string
	DBNamePrefix    string
	Timeout         int
	MaxPoolSize     uint64
	IdleConnTimeout int
}

// GetMongoDBConfig load DB configuration for Mongo Database
// envPrefix define the prefix name (e.g. 'STUDY_') to get configuration for a specific db
func GetMongoDBConfig(envPrefix string) DBConfig {
	connStr := os.Getenv(envPrefix + "DB_CONNECTION_STR")
	username := os.Getenv(envPrefix + "DB_USERNAME")
	password := os.Getenv(envPrefix + "DB_PASSWORD")
	prefix := os.Getenv("DB_CONNECTION_PREFIX") // Used in test mode
	if connStr == "" || username == "" || password == "" {
		logger.Fatal("Couldn't read DB credentials.")
	}

	URI := fmt.Sprintf(`mongodb%s://%s:%s@%s`, prefix, username, password, connStr)

	var err error
	Timeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		Fatalf(" error in `DB_TIMEOUT`: %s", err)
	}
	IdleConnTimeout, err := strconv.Atoi(os.Getenv("DB_IDLE_CONN_TIMEOUT"))
	if err != nil {
		Fatalf("error in `DB_IDLE_CONN_TIMEOUT` : %s", err.Error())
	}
	mps, err := strconv.Atoi(os.Getenv("DB_MAX_POOL_SIZE"))
	MaxPoolSize := uint64(mps)
	if err != nil {
		Fatalf("error in `DB_MAX_POOL_SIZE` : %s", err)
	}

	DBNamePrefix := os.Getenv("DB_DB_NAME_PREFIX")

	return DBConfig{
		URI:             URI,
		Timeout:         Timeout,
		IdleConnTimeout: IdleConnTimeout,
		MaxPoolSize:     MaxPoolSize,
		DBNamePrefix:    DBNamePrefix,
	}
}

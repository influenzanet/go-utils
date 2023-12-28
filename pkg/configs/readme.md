# Configuration Utility package

This package provides several utility functions to handle configuration of Influenzanet services.

### Database config

- GetMongoDBConfig(prefix) : Load MongoDB configuration config with common env name. Returns DBConfig type

### Env utilities

- ParseDuration(string, defaultUnit): Parse Duration string. If number provided without unit, consider it with the defaultUnit
- ParseEnvDuration(): parse duration from en environment variable, with default value if not provided
- RequireEnv(string): get an env value, raise a fatal error if not provided or empty
- GetEnvInt(string, default): Get an integer value from env

### Internal logger

These function can log some step/errors. By default it uses common `log` package function, it's possible to override using SetConfigLogger() function passing a 
ConfigLogger struct

Default is :
ConfigLogger{
	Fatal: log.Fatal,
	Error: log.Println,
	Info:  log.Println,
}

Fatal logger is expected to stop the program after logged.
package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv retrieves the "ENV" environment variable.
func GetEnv() string {
	return getEnvironmentValue("ENV") // ❶
}

// GetDataSourceURL retrieves the "DATA_SOURCE_URL" environment variable.
func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL") // ❷
}

// GetApplicationPort retrieves the "APPLICATION_PORT" and converts it to an integer.
func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT") // ❸
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

// getEnvironmentValue fetches an environment variable and ensures it is set.
func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" { // ❹
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}

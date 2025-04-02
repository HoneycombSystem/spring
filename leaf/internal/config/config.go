package config

import (
	"fmt"
	"os"
)

func getEnv(key string) string {
	// Get the value of the key from the environment
	value := os.Getenv(key)
	if value == "" {
		// If the value is empty, return an error
		return fmt.Errorf("environment variable %s is not set", key)
	}
	return value
}
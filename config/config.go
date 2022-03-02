package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	log.Debugln("Reading .env file for application configuration")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func Debug() bool {
	return Config("DEBUG") == "true"
}

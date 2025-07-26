// Package config handles application configuration, including loading environment variables.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	EnvironmentName string
	DBDialect       string
	DBUsername      string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBDebug         string
	ThrottleTTL     string
	ThrottleLimit   string
}

var AppConfig Configuration

func init() {
	env := os.Getenv("GO_ENV")
	fmt.Println(env)
	if env == "" {
		env = "dev"
	}

	envFile := ".env." + env
	// cwd, _ := os.Getwd()
	// log.Printf("Loading environment file from: %s", cwd)

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig = Configuration{
		EnvironmentName: os.Getenv("ENVIRONMENT_NAME"),
		DBDialect:       os.Getenv("DB_DIALECT"),
		DBUsername:      os.Getenv("DB_USERNAME"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBDebug:         os.Getenv("DB_DEBUG"),
		ThrottleTTL:     os.Getenv("THROTTLE_TTL"),
		ThrottleLimit:   os.Getenv("THROTTLE_LIMIT"),
	}
	log.Println("Loading the", AppConfig.EnvironmentName)
}

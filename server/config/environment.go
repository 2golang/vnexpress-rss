package config

import (
	"log"
	"os"
	"strconv"

	"github.com/2golang/vnexpress-rss/helpers"
	env "github.com/joho/godotenv"
)

type Environment struct {
	Port string
	Env  string
}

func getDefaultEnvFilePath(envFilePath string) string {
	var priorityEnvFiles = []string{".env", ".env.local", ".env.development", ".env.test", ".env.production"}
	if envFilePath != "" {
		priorityEnvFiles = append(priorityEnvFiles, envFilePath)
	}
	for _, envFile := range priorityEnvFiles {
		if helpers.FileExists(envFile) {
			return envFile
		}
	}
	return ""
}

func LoadEnv(envFilePath string) {

	var usedEnvFile string = getDefaultEnvFilePath(envFilePath)
	// skip loading env file if it's not found
	if usedEnvFile == "" {
		return
	}
	err := env.Load(usedEnvFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err != nil {
		log.Fatal("Error loading config file")
	}
}

func GetEnv() Environment {
	return Environment{
		Port: parseEnvParam("PORT", "8080").(string),
		Env:  parseEnvParam("ENV", "development").(string),
	}
}

func parseEnvParam(key string, defaultValue any) any {
	switch v := defaultValue.(type) {
	case string:
		value := os.Getenv(key)
		if value == "" {
			return v
		}
		return value
	case int:
		value := os.Getenv(key)
		if value == "" {
			return v
		}
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return parsedValue
	case bool:
		value := os.Getenv(key)
		if value == "" {
			return v
		}
		parsedValue, err := strconv.ParseBool(value)
		if err != nil {
			panic(err)
		}
		return parsedValue
	case float64:
		value := os.Getenv(key)
		if value == "" {
			return v
		}
		parsedValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		return parsedValue
	default:
		value := os.Getenv(key)
		if value == "" {
			return v
		}
		return value
	}

}

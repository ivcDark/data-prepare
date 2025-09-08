package config

import (
	"os"
)

type Config struct {
	RabbitMQURL string
	DatabaseURL string
}

func Load() (*Config, error) {
	return &Config{
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://user:password@rabbitmq:5672/"),
		DatabaseURL: getEnv("DATABASE_URL", "user:password@tcp(mysql:3306)/football"),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
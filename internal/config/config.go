package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPPort        int
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {
	port, err := intFromEnv("HTTP_PORT", 8080)
	if err != nil {
		return nil, err
	}

	return &Config{
		HTTPPort:        port,
		ShutdownTimeout: 10 * time.Second,
	}, nil
}

func (c *Config) HTTPAddr() string {
	return fmt.Sprintf(":%d", c.HTTPPort)
}

func intFromEnv(key string, fallback int) (int, error) {
	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("config: %s must be an integer, got %q", key, raw)
	}

	return value, nil
}

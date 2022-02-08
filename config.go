package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	URL         string
	Env         string
	AppSecret   string
	DatabaseURL string
	Email       EmailConfig
}

func LoadConfig() (Config, error) {
	emailConfig := EmailConfig{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		FromEmail:    os.Getenv("FROM_EMAIL"),
	}

	config := Config{
		URL:         os.Getenv("APP_URL"),
		Env:         os.Getenv("APP_ENV"),
		AppSecret:   os.Getenv("APP_SECRET"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Email:       emailConfig,
	}

	if emailConfig.SMTPHost == "" {
		return config, errors.New("must provide SMTP_HOST")
	}
	if emailConfig.SMTPUsername == "" {
		return config, errors.New("must provide SMTP_USERNAME")
	}
	if emailConfig.SMTPPassword == "" {
		return config, errors.New("must provide SMTP_PASSWORD")
	}
	if emailConfig.FromEmail == "" {
		return config, errors.New("must provide FROM_EMAIL")
	}
	if config.URL == "" {
		config.URL = "http://localhost:8080"
	}
	if config.Env == "" {
		config.Env = "debug"
	}
	if config.AppSecret == "" {
		return config, errors.New("must provide APP_SECRET")
	}
	if config.DatabaseURL == "" {
		return config, errors.New("must provide DATABASE_URL")
	}
	if !strings.Contains(config.DatabaseURL, "sslmode=disable") {
		config.DatabaseURL = fmt.Sprintf("%s?sslmode=disable", config.DatabaseURL)
	}

	return config, nil
}

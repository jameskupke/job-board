package main

import (
	"errors"
	"os"
)

type Config struct {
	Email       EmailConfig
	DatabaseURL string
	AppSecret   string
	Env         string
}

func LoadConfig() (Config, error) {
	emailConfig := EmailConfig{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		FromEmail:    os.Getenv("FROM_EMAIL"),
	}

	config := Config{
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
	if config.Env == "" {
		config.Env = "debug"
	}
	if config.AppSecret == "" {
		return config, errors.New("must provide APP_SECRET")
	}
	if config.DatabaseURL == "" {
		return config, errors.New("must provide DATABASE_URL")
	}

	return config, nil
}

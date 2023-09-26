package apiserver

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BindAddr string 
	LogLevel string 
	DatabaseURL string
}

func NewConfig() *Config {
	err := godotenv.Load()
  if err != nil {
    fmt.Printf("Error loading .env file")
  }
	return &Config{
		BindAddr: os.Getenv("BINDADDR"),
		LogLevel: os.Getenv("LOGLEVEL"),
		DatabaseURL: os.Getenv("DBURL"),
	}
}
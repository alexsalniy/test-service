package main

import (
	"github.com/joho/godotenv"
	"flag"
	"log"

	"github.com/alexsalniy/test-service/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()


	config := apiserver.NewConfig()
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
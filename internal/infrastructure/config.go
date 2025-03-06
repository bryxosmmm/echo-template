package infrastructure

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Address      string `env:"SERVER_ADDRESS"`
		AllowedHosts []string
	}
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
	}
	Other struct {
		JWTKey string `env:"JWT"`
	}
	Debug bool
}

func LoadConfig() *Config {
	godotenv.Load()
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatal("can't get config")
		return nil
	}
	config.Debug = slices.Contains(
		[]string{"true", "1", "on", "yes"},
		strings.ToLower(os.Getenv("DEBUG")),
	)
	config.Server.AllowedHosts = strings.Split(os.Getenv("ALLOWED_HOSTS"), ";")
	return &config
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Addr  string
}

func LoadCfg() *Config {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	cfg := &Config{
		DBUrl: dbUrl,
		Addr:  ":8080",
	}
	return cfg

}

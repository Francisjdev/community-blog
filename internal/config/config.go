package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl     string
	Addr      string
	SecretKey string
}

func LoadCfg() *Config {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	sKey := os.Getenv("SECRET_KEY")
	cfg := &Config{
		DBUrl:     dbUrl,
		Addr:      ":8080",
		SecretKey: sKey,
	}
	return cfg

}

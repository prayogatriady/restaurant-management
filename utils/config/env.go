package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Port       string
		Mode       string
		AppName    string
		AppVersion string
		Timezone   string
	}

	Http struct {
		Secret string
	}

	Db struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

var AppCfg *AppConfig

func InitEnv() {

	var config AppConfig

	config.App.Port = GetEnv("PORT", "8000")
	config.App.Mode = GetEnv("MODE", "local")
	config.App.AppName = GetEnv("APP_NAME", "Restaurant ")
	config.App.AppVersion = GetEnv("APP_VERSION", "1.0.0")
	config.App.Timezone = GetEnv("TIMEZONE", "Asia/Jakarta")
	config.Http.Secret = GetEnv("SECRET", "")

	config.Db.User = GetEnv("MYSQL_USER", "")
	config.Db.Password = GetEnv("MYSQL_PASSWORD", "")
	config.Db.Host = GetEnv("MYSQL_HOST", "")
	config.Db.Port = GetEnv("MYSQL_PORT", "")
	config.Db.Name = GetEnv("MYSQL_NAME", "")

	AppCfg = &config
}

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

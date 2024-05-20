package config

import (
	"os"

	"github.com/joho/godotenv"
)

type HttpServer struct {
	Addr string `json:"addr"`
}

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Env struct {
	HttpServer   HttpServer   `json:"httpServer"`
	Application  Application  `json:"application"`
	Environment  string       `json:"environment"`
	DbConnection DbConnection `json:"dbConnection"`
}

type DbConnection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

func getEnv(keyValue string, defaultValue string) string {
	if env, ok := os.LookupEnv(keyValue); ok && len(env) > 0 {
		return env
	}
	return defaultValue
}

func LoadEnv() Env {
	godotenv.Load()

	httpServer := HttpServer{
		Addr: getEnv("HTTP_ADDR", ":8080"),
	}

	application := Application{
		Name:    getEnv("APP_NAME", "local_app"),
		Version: getEnv("APP_VERSION", "1.0.0-local"),
	}

	dbConnection := DbConnection{
		User:     getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
	}

	env := Env{
		HttpServer:   httpServer,
		Application:  application,
		Environment:  getEnv("ENV", "local"),
		DbConnection: dbConnection,
	}

	return env
}

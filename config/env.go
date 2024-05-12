package config

import "os"

type HttpServer struct {
	Addr string `json:"addr"`
}

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Env struct {
	HttpServer  HttpServer  `json:"httpServer"`
	Application Application `json:"application"`
	Environment string      `json:"environment"`
}

func getEnv(keyValue string, defaultValue string) string {
	if env, ok := os.LookupEnv(keyValue); ok && len(env) > 0 {
		return env
	}
	return defaultValue
}

func LoadEnv() Env {
	httpServer := HttpServer{
		Addr: getEnv("HTTP_ADDR", ":8080"),
	}

	application := Application{
		Name:    getEnv("APP_NAME", "local_app"),
		Version: getEnv("APP_VERSION", "1.0.0-local"),
	}

	env := Env{
		HttpServer:  httpServer,
		Application: application,
		Environment: getEnv("ENV", "local"),
	}

	return env
}

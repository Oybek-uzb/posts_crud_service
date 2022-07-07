package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostsServiceHost string
	PostsServicePort int

	HttpPort string
}

func Load() Config {
	config := Config{}

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8081"))

	config.PostsServiceHost = cast.ToString(getOrReturnDefault("POSTS_SERVICE_HOST", "localhost"))
	config.PostsServicePort = cast.ToInt(getOrReturnDefault("POSTS_SERVICE_PORT", 8082))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

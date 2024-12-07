package config

type Config struct {
	AppID string
}

func NewConfig(getenv func(string) string) Config {
	return Config{
		AppID: getenv("APP_ID"),
	}
}

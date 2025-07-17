package config

type Config struct {
	BaseURL string `json:"base_url"`
}

func New() *Config {
	return &Config{
		BaseURL: "http://localhost:3000/api/v1/students",
	}
}
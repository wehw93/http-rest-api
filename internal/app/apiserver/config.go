package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml: "log_level"`
	dataBaseURL string `toml: "database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "localhost:8080",
		LogLevel:    "debug",
		dataBaseURL: "host=localhost dbname=restapi_dev user=postgres password=pwd123 sslmode=disable",
	}
}

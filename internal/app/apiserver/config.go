package apiserver

// Config
type Config struct {
	// Server address
	BindAddr string `toml:"bind_addr"`
	// Logger level
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// NewConfig - initialisation with default parameteres
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

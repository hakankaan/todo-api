package config

// Config is a struct for configuration
type Config struct {
	Http
	Postgres
	Redis
	Logging
}

// Http is a struct for http configuration
type Http struct {
	Port string `yaml:"port"`
}

// Postgres is a struct for postgres configuration
type Postgres struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
	DB   string `yaml:"database"`
	SSL  string `yaml:"sslmode"`
}

// Redis is a struct for redis configuration
type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Logging is a struct for logging configuration
type Logging struct {
	Level string `yaml:"level"`
}

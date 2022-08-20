package config

type Config struct {
	Filenames []string
}

func New(filenames []string) *Config {
	return &Config{Filenames: filenames}
}

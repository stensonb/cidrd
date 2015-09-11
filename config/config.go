package config

type Config struct {
  Port uint16
}

func GetConfig() *Config {
  return &Config{Port: 8080}
}

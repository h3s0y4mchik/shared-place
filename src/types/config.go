package types

type Config struct {
	addr string
	port uint
}

func NewConfig(addr string, port uint) (*Config, error) {
	return &Config{}, nil
}

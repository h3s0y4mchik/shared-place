package types

type conn int

type Server struct {
	cfg *Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		cfg,
	}, nil
}

func (s *Server) validateAddr() error {
	return nil
}

func (s *Server) connect() (*conn, error) {
	c := new(conn)
	*c = 0
	return c, nil
}

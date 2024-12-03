package types

type conn int

type Server struct {
	addr string
	port uint
	dir  string
}

func NewServer(addr string, port uint, dir string) (*Server, error) {
	if err := validateAddr(addr); err != nil {
		return nil, err
	}
	return &Server{
		addr, port, dir,
	}, nil
}

func validateAddr(addr string) error {
	_ = addr
	return nil
}

func (s *Server) Connect() (*conn, error) {
	c := new(conn)
	*c = 0
	return c, nil
}

package gotur

type Server struct {
	protocol string
	socket   Socket
}

func NewServer(protocol string, socket Socket) *Server {
	return &Server{
		protocol: protocol,
		socket:   socket,
	}
}

func (s *Server) Start(address string) error {
	if err := s.socket.Bind(address); err != nil {
		return err
	}
	return s.socket.Listen()
}
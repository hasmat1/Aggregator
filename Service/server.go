package Service

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s *Server) mustEmbedUnimplementedEndpointsServer() {
	panic("implement me")
}

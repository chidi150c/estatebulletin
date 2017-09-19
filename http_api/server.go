//Server: ROle is to open the socket and serve the handler
//HTTP server
package http_api

type Server struct{
	ln net.Listener

	//Handler to serve
	Handler *Handler
	
	//Bind address (port) to open
	Address string
}

func NewServer(port string, h *Handler)*Server{
return &Server{ 
	Address : ":" + port,
	Handler : h,
}
}

//Open opens a socket and serves the HTTP server
func (s *Server) Open() error{
	//Open socket
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil{
		return err
	}
	s.ln = ln
}

//Start HTTP server
go func() {http.Serve(s.ln, s.Handler)}()

return nil

//Close xloses the socjet
func (s *Server) Close() error{
	if s.ln != nil{
		s.ln.Close()
	}
	return nil
}

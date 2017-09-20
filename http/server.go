//Server: ROle is to open the socket and serve the handler
//HTTP server
package http

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type Server struct {
	ln net.Listener

	//Handler to serve
	Handler *HttpHandler

	//Bind address (port) to open
	Address string
}

func NewServer(port string, h *HttpHandler) *Server {
	return &Server{
		Address: ":" + port,
		Handler: h,
	}
}

//Open opens a socket and serves the HTTP server
func (s *Server) Open(done chan bool, sigs chan os.Signal) error {
	//Open socket
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	s.ln = ln
	// Start HTTP server.
	go func() {

		log.Fatal(http.Serve(s.ln, handlers.CombinedLoggingHandler(os.Stderr, s.Handler)))
		defer s.ln.Close()
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	return nil
}

//Close closes the socket
func (s *Server) Close() {
	if s.ln != nil {
		s.ln.Close()
	}
}

// Port returns the port that the server is open on. Only valid after open.
func (s *Server) Port() int {
	return s.ln.Addr().(*net.TCPAddr).Port
}

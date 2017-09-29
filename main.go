package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/appframe/chatting"
	"github.com/appframe/http"
	"google.golang.org/grpc"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	port := os.Getenv("PORT")
	cport := os.Getenv("CPORT")
	appHandler := http.NewAppHandler()
	httpHandler := http.NewHttpHandler(appHandler)
	server := http.NewServer(port, httpHandler)

	gs := grpc.NewServer()
	chatting.RegisterChatServiceServer(gs, &chatting.ChatServer{})
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", cport))
	if err != nil {
		log.Fatalf("Could not listen on Port: %s %v", cport, err)
	}
	go func() {
		fmt.Println("Listening on: ", ln.Addr().(*net.TCPAddr).Port)
		err := gs.Serve(ln)
		if err != nil {
			log.Fatalf("Could not Serve Chat: %v", err)
		}
	}()
	if err := server.Open(done, sigs); err != nil {
		log.Fatalf("Unable to Open Server for listen and serve: %v", err)
	}
	fmt.Println("Listening on: ", server.Port())
	<-done
	fmt.Println("Good Bye")
}

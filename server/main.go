package main

import (
	"google.golang.org/grpc"
	"fmt"
	"github.com/aslanbekirov/grpc-options-test/proto"
	"net"
	"log"
	"golang.org/x/net/context"
)

func main() {
	srv := grpc.NewServer()

	var greetings GreetingServer;
	hello.RegisterGreetingServer(srv, greetings)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("could not listen at 8080:, %v", err)
	}

	log.Println("Listening at :8080")

	log.Fatal(srv.Serve(l))
}

type GreetingServer struct {
}

func (g GreetingServer) SayHello(ctx context.Context, lang *hello.Lang) (*hello.Void, error){
	fmt.Println("hello qaqa")
	return &hello.Void{}, nil
}


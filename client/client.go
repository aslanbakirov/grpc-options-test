package main

import (
	"google.golang.org/grpc"
	"fmt"
	"os"
	"github.com/aslanbekirov/grpc-options-test/proto"
	"context"
)

func main(){
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := hello.NewGreetingClient(conn)
	//SayHello(context.Background(), &hello.Lang{}, &client)
	client.SayHello(context.Background(), &hello.Lang{})
}




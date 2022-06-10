package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Id        = "id"
	FirstName = "first-name"
	LastName  = "last-name"
	Nickname  = "nickname"
	Password  = "password"
	Email     = "email"
	Country   = "country"
)

// NewGRPCConn returns a gRPC client connection.
func NewGRPCConn(host, port string) *grpc.ClientConn {
	conn, err := grpc.Dial(net.JoinHostPort(host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect: ", err)
	}
	return conn
}

// ResponsePrinter marshals to JSON and then prints to stdout
func ResponsePrinter(r interface{}) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal("error formatting create response: ", err.Error())
	}
	_, _ = fmt.Fprintf(os.Stdout, "%s", b)
}

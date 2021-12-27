package main

import (
	"log"
	"net"
	"server-grpc/application/services"
	"server-grpc/config"
	"server-grpc/grpc/model"
	"server-grpc/grpc/service"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	model.RegisterMovieServicesServer(srv, &service.MovieServer{
		MovieServices: services.SetupMovieServices(),
	})

	log.Println("Starting RPC server at", config.GrpcPort())

	l, err := net.Listen("tcp", config.GrpcPort())
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GrpcPort(), err)
	}

	log.Fatal(srv.Serve(l))
}

package main

import (
	"log"
	"net"

	"github.com/himanshuk42/grpcpokemon/pkg/model"
	pb "github.com/himanshuk42/grpcpokemon/proto"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:50051"

func main() {
	model.Init()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("couldn't listen on address: %v\n", err)
	}

	log.Printf("listening on address: %v\n", address)

	s := grpc.NewServer()
	pb.RegisterPokemonServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("couldn't server: %v\n", err)
	}
}

package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/himanshuk42/grpcpokemon/proto"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to the gRPC Server: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPokemonServiceClient(conn)

	// data := &model.PokemonItem{
	// 	Pid:         "3",
	// 	Name:        "test",
	// 	Power:       "50",
	// 	Description: "test sample for testing",
	// }
	// updatedData := model.PokemonItem{
	// 	Pid:         "3",
	// 	Name:        "hello",
	// 	Power:       "100",
	// 	Description: "converted to hello",
	// }

	// createPokemon(c, data)
	// _ := createPokemon(c, data)
	id := "2"
	readPokemon(c, id)
	// updatePokemon(c, updatedData)
	// deletePokemon(c, id)
	// readAllPokemon(c)
}

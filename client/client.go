package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/himanshuk42/grpcpokemon/model"
	pb "github.com/himanshuk42/grpcpokemon/proto"
)

func createPokemon(c pb.PokemonServiceClient, payload *model.PokemonItem) string {
	fmt.Println("createPokemon function calleds")
	req := &pb.CreatePokemonRequest{
		Pokemon: &pb.Pokemon{
			Pid:         payload.Pid,
			Name:        payload.Name,
			Power:       payload.Power,
			Description: payload.Description,
		},
	}
	fmt.Println(req)

	response, err := c.CreatePokemon(context.Background(), req)
	if err != nil {
		log.Fatalf("error while creating pokemon: %v\n", err)
	}

	fmt.Println("Succesfully created the pokemon")
	fmt.Println(response.Pokemon)
	return response.Pokemon.Pid
}

func readPokemon(c pb.PokemonServiceClient, pid string) {
	req := &pb.ReadPokemonRequest{
		Pid: pid,
	}
	response, err := c.ReadPokemon(context.Background(), req)
	if err != nil {
		log.Fatalf("error while reading pokemon from client: %v\n", err)
	}
	fmt.Println("Successfully retrieved the pokemon")
	fmt.Println(response.Pokemon)
}

func updatePokemon(c pb.PokemonServiceClient, payload model.PokemonItem) {
	req := &pb.UpdatePokemonRequest{
		Pokemon: &pb.Pokemon{
			Pid:         payload.Pid,
			Name:        payload.Name,
			Power:       payload.Power,
			Description: payload.Description,
		},
	}

	response, err := c.UpdatePokemon(context.Background(), req)
	if err != nil {
		log.Fatalf("error while updating pokemon from client: %v\n", err)
	}

	fmt.Println("Successfully updated the pokemon")
	fmt.Println(response.Pokemon)
}
func deletePokemon(c pb.PokemonServiceClient, pid string) {
	req := &pb.DeletePokemonRequest{
		Pid: pid,
	}
	response, err := c.DeletePokemon(context.Background(), req)
	if err != nil {
		log.Fatalf("error while deleting pokemon from client: %v\n", err)
	}

	fmt.Println("Successfully deleted the pokemon")
	fmt.Println(response.Pokemon)
}

func readAllPokemon(c pb.PokemonServiceClient) {
	fmt.Println("readAllPokemon function was invoked")

	req := &pb.ListPokemonRequest{}
	stream, err := c.ListPokemon(context.Background(), req)
	if err != nil {
		log.Fatalf("couldn't list out the pokemon: %v\n", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receving stream of response: %v\n", err)
		}

		fmt.Println(response.Pokemon)
	}
}

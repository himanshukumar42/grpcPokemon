package main

import (
	"context"
	"fmt"

	"github.com/himanshuk42/grpcpokemon/pkg/model"
	pb "github.com/himanshuk42/grpcpokemon/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.PokemonServiceServer
}

func (s *Server) CreatePokemon(ctx context.Context, req *pb.CreatePokemonRequest) (*pb.CreatePokemonResponse, error) {
	fmt.Println("CreatePokemon function was invoked")
	data := req.GetPokemon()
	pokemon := &model.PokemonItem{
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}

	fmt.Println("======reached here")
	pokemon, err := pokemon.CreatePokemon()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}

	fmt.Println("======reached here")

	return &pb.CreatePokemonResponse{
		Pokemon: &pb.Pokemon{
			Pid:         pokemon.Pid,
			Name:        pokemon.Name,
			Power:       pokemon.Power,
			Description: pokemon.Description,
		},
	}, nil
}

func (s *Server) ReadPokemon(ctx context.Context, req *pb.ReadPokemonRequest) (*pb.ReadPokemonResponse, error) {
	fmt.Println("ReadPokemon function was invoked")
	pid := req.Pid

	pokemon := &model.PokemonItem{}
	pokemon, err := pokemon.FindPokemonByID(pid)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("couldn't read pokemon: %v\n", err),
		)
	}

	return &pb.ReadPokemonResponse{
		Pokemon: &pb.Pokemon{
			Pid:         pokemon.Pid,
			Name:        pokemon.Name,
			Power:       pokemon.Power,
			Description: pokemon.Description,
		},
	}, nil
}

func (s *Server) UpdatePokemon(ctx context.Context, req *pb.UpdatePokemonRequest) (*pb.UpdatePokemonResponse, error) {
	fmt.Println("UpdatePokemon function was invoked")
	data := req.GetPokemon()
	pokemon := &model.PokemonItem{
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}

	pokemon, err := pokemon.UpdatePokemon(pokemon.Pid)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Couldn't update pokemon: %v\n", err),
		)
	}

	return &pb.UpdatePokemonResponse{
		Pokemon: &pb.Pokemon{
			Pid:         pokemon.Pid,
			Name:        pokemon.Name,
			Power:       pokemon.Power,
			Description: pokemon.Description,
		},
	}, nil
}

func (s *Server) DeletePokemon(ctx context.Context, req *pb.DeletePokemonRequest) (*pb.UpdatePokemonResponse, error) {
	fmt.Println("DeletePokemon function was invoked")
	pid := req.Pid

	pokemon := &model.PokemonItem{}
	_, err := pokemon.DeletePokemon(pid)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Couldn't delete pokemon: %v\n", err),
		)
	}

	return &pb.UpdatePokemonResponse{}, nil
}

func (s *Server) ListPokemon(req *pb.ListPokemonRequest, stream pb.PokemonService_ListPokemonServer) error {
	fmt.Println("ListPokemon function was invoked")

	pokemon := &model.PokemonItem{}
	pokemons, err := pokemon.FindAllPokemon()
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("couldn't list pokemon: %v\n", err),
		)
	}

	for _, pokemon := range *pokemons {
		err := stream.Send(&pb.ListPokemonResponse{
			Pokemon: &pb.Pokemon{
				Pid:         pokemon.Pid,
				Name:        pokemon.Name,
				Power:       pokemon.Power,
				Description: pokemon.Description,
			},
		})
		if err != nil {
			return status.Errorf(
				codes.NotFound,
				fmt.Sprintf("couldn't stream pokemon: %v\n", err),
			)
		}
	}
	return nil
}

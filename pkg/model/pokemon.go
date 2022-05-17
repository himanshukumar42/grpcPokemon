package model

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/himanshuk42/grpcpokemon/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

type PokemonItem struct {
	gorm.Model
	Pid         string `json:"pid"`
	Name        string `json:"name"`
	Power       string `json:"power"`
	Description string `json:"description"`
}

func Init() {
	// connectionString := utils.GetConnectionString()
	database, err := gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}

	db = database
	db.AutoMigrate(PokemonItem{})
}

func getPokemonData(data *PokemonItem) *pb.Pokemon {
	return &pb.Pokemon{
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}
}

func (pokemon *PokemonItem) CreatePokemon() (*PokemonItem, error) {
	err := db.Debug().Create(&pokemon).Error
	if err != nil {
		return &PokemonItem{}, err
	}
	return pokemon, nil
}

func (pokemon *PokemonItem) FindAllPokemon() (*[]PokemonItem, error) {
	var err error
	pokemons := []PokemonItem{}
	err = db.Debug().Model(&PokemonItem{}).Limit(100).Find(&pokemons).Error
	if err != nil {
		return &[]PokemonItem{}, err
	}

	fmt.Println(pokemons)
	return &pokemons, nil
}

func (pokemon *PokemonItem) FindPokemonByID(pid string) (*PokemonItem, error) {
	err := db.Debug().Model(PokemonItem{}).Where("id = ?", pid).Take(&pokemon).Error
	if err != nil {
		return &PokemonItem{}, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return &PokemonItem{}, errors.New("pokemon not found")
	}
	return pokemon, nil
}

func (pokemon *PokemonItem) UpdatePokemon(pid string) (*PokemonItem, error) {
	db = db.Debug().Model(&PokemonItem{}).Where("id = ?", pid).Take(&PokemonItem{}).UpdateColumns(
		map[string]interface{}{
			"name":        pokemon.Name,
			"power":       pokemon.Power,
			"description": pokemon.Description,
		},
	)

	if db.Error != nil {
		return &PokemonItem{}, db.Error
	}

	err := db.Debug().Model(&PokemonItem{}).Where("id = ? ", pid).Take(&PokemonItem{}).Error
	if err != nil {
		return &PokemonItem{}, nil
	}
	return pokemon, nil
}

func (pokemon *PokemonItem) DeletePokemon(pid string) (int64, error) {
	db = db.Debug().Model(&PokemonItem{}).Where("id = ?", pid).Take(&PokemonItem{}).Delete(&PokemonItem{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

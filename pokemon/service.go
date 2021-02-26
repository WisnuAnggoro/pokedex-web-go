package pokemon

import (
	"fmt"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type service struct {
	pokemonSprite string
}

type Service interface {
	GetPokemonCardList(int, int) (*[]PokemonCard, error)
	// GetPokemons() (*structs.Resource, error)
	GetPokemonByID(string) (*structs.Pokemon, error)
	// GetPokemonSpeciesByID(string) (*structs.PokemonSpecies, error)
}

func NewService(pokemonSprite string) Service {
	return &service{
		pokemonSprite: pokemonSprite,
	}
}

func (s *service) GetPokemonCardList(offset, limit int) (*[]PokemonCard, error) {
	ret := []PokemonCard{}
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	i := 0
	for i < limit {
		i++
		p, err := s.GetPokemonByID(strconv.Itoa(offset + i))

		if err != nil {
			return nil, err
		}

		types := []string{}
		for _, v := range p.Types {
			types = append(types, v.Type.Name)
		}
		d := PokemonCard{
			ID:     p.ID,
			Name:   p.Name,
			Sprite: fmt.Sprintf(s.pokemonSprite, p.ID),
			Types:  types,
		}
		ret = append(ret, d)
	}

	return &ret, nil
}

func (s *service) GetPokemonByID(id string) (*structs.Pokemon, error) {
	p, err := pokeapi.Pokemon(id)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

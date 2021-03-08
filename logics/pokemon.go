package logics

import (
	"fmt"
	"net/http"

	"github.com/wisnuanggoro/pokedex-web-go/config"
)

type pokemonLogic struct {
	cfg config.Config
}

type PokemonLogic interface {
	FormatID(id int) string
	GetSpriteURLByID(id int) string
}

func NewPokemonLogic(cfg config.Config) PokemonLogic {
	return &pokemonLogic{
		cfg: cfg,
	}
}

func (l *pokemonLogic) FormatID(id int) string {
	return fmt.Sprintf("%03d", id)
}

func (l *pokemonLogic) GetSpriteURLByID(id int) string {
	sprites := l.cfg.PokemonSprites
	pokeSprite := ""
	for _, sprite := range sprites {
		checkingSprite := fmt.Sprintf(sprite, id)
		resp, err := http.Head(checkingSprite)
		if err == nil && resp.StatusCode == http.StatusOK {
			pokeSprite = checkingSprite
			break
		}
	}
	if pokeSprite == "" {
		pokeSprite = "/assets/img/pokeball.png"
	}

	return pokeSprite
}

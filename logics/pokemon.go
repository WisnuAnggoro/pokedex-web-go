package logics

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/models"
)

type pokemonLogic struct {
	cfg config.Config
}

type PokemonLogic interface {
	FormatID(id int) string
	GetSpriteURLByID(id int) string
	CreatePokemonCardList(results []structs.Result) []models.PokemonCard
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

func (l *pokemonLogic) CreatePokemonCardList(results []structs.Result) []models.PokemonCard {
	var cardList []models.PokemonCard
	totalData := len(results)
	if totalData > 0 {
		var wg sync.WaitGroup
		wg.Add(totalData)

		for _, v := range results {
			go func(v structs.Result) {
				defer wg.Done()

				// Get Search Term
				// (Search by ID is preferred than by Name)
				searchTerm := v.Name
				splitFn := func(c rune) bool {
					return c == '/'
				}
				token := strings.FieldsFunc(v.URL, splitFn)
				tokenLen := len(token)
				if tokenLen > 0 {
					searchTerm = token[tokenLen-1]
				}

				// Get Pokemon Detail
				pokeDetail, err := pokeapi.Pokemon(searchTerm)
				if err != nil {
					log.Errorf("Error in getting detail for '%s': %s", searchTerm, err.Error())
				}

				// Get Pokemon Sprite
				pokeID := pokeDetail.ID
				pokeSprite := ""
				for _, sprite := range l.cfg.PokemonSprites {
					checkingSprite := fmt.Sprintf(sprite, pokeID)
					resp, err := http.Head(checkingSprite)
					if err == nil && resp.StatusCode == http.StatusOK {
						pokeSprite = checkingSprite
						break
					}
				}
				if pokeSprite == "" {
					pokeSprite = "/assets/img/pokeball.png"
				}

				// Map Pokemon Types
				pokeTypes := []string{}
				for _, t := range pokeDetail.Types {
					pokeTypes = append(pokeTypes, t.Type.Name)
				}

				// Construct Pokemon Card
				pokemonCard := models.PokemonCard{
					ID:          pokeID,
					IDFormatted: fmt.Sprintf("%03d", pokeID),
					Name:        pokeDetail.Name,
					Sprite:      pokeSprite,
					Types:       pokeTypes,
				}

				// Append to Pokemon Card List
				cardList = append(cardList, pokemonCard)
			}(v)
		}

		wg.Wait()
	}

	// Sort cardList by ID ascending
	sort.Slice(cardList[:], func(i, j int) bool {
		if cardList[i].ID < cardList[j].ID {
			return true
		}
		return false
	})

	return cardList
}

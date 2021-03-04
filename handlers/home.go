package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/models"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type handler struct {
	cfg          config.Config
	render       render.Render
	errorHandler ErrorHandler
}

type HomeHandler interface {
	HomePage(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(cfg config.Config, render render.Render, errorHandler ErrorHandler) HomeHandler {
	return &handler{
		cfg:          cfg,
		render:       render,
		errorHandler: errorHandler,
	}
}

func (h *handler) HomePage(w http.ResponseWriter, r *http.Request) {
	limit := 96
	offset := 0

	urlParams := r.URL.Query()
	page, ok := urlParams["page"]
	if ok {
		p, err := strconv.Atoi(page[0])
		if err == nil && p > 0 {
			offset = (p - 1) * limit
		}
	}

	pokemon, err := pokeapi.Resource("pokemon", offset, limit)
	if err != nil {
		log.Errorf("Error in getting pokemon resource: %s", err.Error())
		return
	}

	// Create Card List
	pokemonCardList := CreatePokemonCardList(pokemon.Results, h.cfg.PokemonSprites)

	// Setup Pagination
	currentPage := offset/limit + 1
	totalPage := pokemon.Count/limit + 1
	pageList := make([]int, totalPage)
	for i := 0; i < totalPage; i++ {
		pageList[i] = i + 1
	}
	pagination := models.Pagination{
		PreviousPage: currentPage - 1,
		CurrentPage:  currentPage,
		NextPage:     currentPage + 1,
		TotalPage:    totalPage,
		PageList:     pageList,
	}

	// Send data to template
	data := make(map[string]interface{})
	data["pokemon_card_list"] = pokemonCardList
	data["pagination"] = pagination
	h.render.RenderTemplate(w, "page.home.gohtml", &models.TemplateData{
		Data: data,
	})

}

func CreatePokemonCardList(results []structs.Result, sprites []string) []models.PokemonCard {
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
				for _, sprite := range sprites {
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

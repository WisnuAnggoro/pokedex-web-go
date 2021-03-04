package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
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
	CardList(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(cfg config.Config, render render.Render, errorHandler ErrorHandler) HomeHandler {
	return &handler{
		cfg:          cfg,
		render:       render,
		errorHandler: errorHandler,
	}
}

func (h *handler) CardList(w http.ResponseWriter, r *http.Request) {
	limit := 48
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
		fmt.Errorf("Error in getting pokemon resource: %s", err.Error())
		return
	}

	// Create Card List
	pokemonCardList := []models.PokemonCard{}
	for _, v := range pokemon.Results {
		pokeDetail, err := pokeapi.Pokemon(v.Name)
		if err != nil {
			fmt.Errorf("Error in getting detail for '%s': %s", v.Name, err.Error())
			return
		}

		pokeTypes := []string{}
		for _, t := range pokeDetail.Types {
			pokeTypes = append(pokeTypes, t.Type.Name)
		}

		pokemonCard := models.PokemonCard{
			ID:          pokeDetail.ID,
			IDFormatted: fmt.Sprintf("%03d", pokeDetail.ID),
			Name:        pokeDetail.Name,
			Sprite:      fmt.Sprintf(h.cfg.PokemonSprite, pokeDetail.ID),
			Types:       pokeTypes,
		}

		pokemonCardList = append(pokemonCardList, pokemonCard)
	}

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
	fmt.Println(limit, offset, pagination)

	// Send data to template
	data := make(map[string]interface{})
	data["pokemon_card_list"] = pokemonCardList
	data["pagination"] = pagination
	h.render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{
		Data: data,
	})

}

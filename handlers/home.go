package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/logics"
	"github.com/wisnuanggoro/pokedex-web-go/models"
	"github.com/wisnuanggoro/pokedex-web-go/utils/pagination"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type handler struct {
	render       render.Render
	pokemonLogic logics.PokemonLogic
	pagination   pagination.Pagination
	errorHandler ErrorHandler
}

type HomeHandler interface {
	HomePage(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(render render.Render, pokemonLogic logics.PokemonLogic, pagination pagination.Pagination, errorHandler ErrorHandler) HomeHandler {
	return &handler{
		render:       render,
		pokemonLogic: pokemonLogic,
		pagination:   pagination,
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
	pokemonCardList := h.pokemonLogic.CreatePokemonCardList(pokemon.Results)

	// Setup Pagination
	paginationData := h.pagination.GetPagination(offset, limit, pokemon.Count)

	// Send data to template
	data := make(map[string]interface{})
	data["pokemon_card_list"] = pokemonCardList
	data["pagination"] = paginationData
	h.render.RenderTemplate(w, "page.home.gohtml", &models.TemplateData{
		Data: data,
	})
}

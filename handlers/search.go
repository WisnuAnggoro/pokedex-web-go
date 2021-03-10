package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/logics"
	"github.com/wisnuanggoro/pokedex-web-go/models"
	"github.com/wisnuanggoro/pokedex-web-go/utils/pagination"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type searchHandler struct {
	render       render.Render
	pokemonLogic logics.PokemonLogic
	pagination   pagination.Pagination
	errorHandler ErrorHandler
}

type SearchHandler interface {
	SearchPage(w http.ResponseWriter, r *http.Request)
}

func NewSearchHandler(render render.Render, pokemonLogic logics.PokemonLogic, pagination pagination.Pagination, errorHandler ErrorHandler) SearchHandler {
	return &searchHandler{
		render:       render,
		pokemonLogic: pokemonLogic,
		pagination:   pagination,
		errorHandler: errorHandler,
	}
}

func (h *searchHandler) SearchPage(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	q, ok := urlParams["q"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pokeSearch, err := pokeapi.Search("pokemon", q[0])
	if err != nil {
		log.Errorf("Error in getting pokemon detail: %s %s", q[0], err.Error())
		return
	}

	// Create Card List
	pokemonCardList := h.pokemonLogic.CreatePokemonCardList(pokeSearch.Results)

	// Setup Pagination
	paginationData := h.pagination.GetPagination(0, pokeSearch.Count, 0)

	// Send data to template
	data := make(map[string]interface{})
	data["pokemon_card_list"] = pokemonCardList
	data["pagination"] = paginationData
	h.render.RenderTemplate(w, "page.search.gohtml", &models.TemplateData{
		Data: data,
	})
}

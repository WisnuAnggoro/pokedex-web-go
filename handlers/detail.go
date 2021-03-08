package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/logics"
	"github.com/wisnuanggoro/pokedex-web-go/models"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type detailHandler struct {
	render       render.Render
	errorHandler ErrorHandler
	pokemonLogic logics.PokemonLogic
}

type DetailHandler interface {
	DetailPage(w http.ResponseWriter, r *http.Request)
}

func NewDetailHandler(render render.Render, errorHandler ErrorHandler, pokemonLogic logics.PokemonLogic) DetailHandler {
	return &detailHandler{
		render:       render,
		errorHandler: errorHandler,
		pokemonLogic: pokemonLogic,
	}
}

func (h *detailHandler) DetailPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		id = "pikachu"
	}

	pokeDetail, err := pokeapi.Pokemon(id)
	if err != nil {
		log.Errorf("Error in getting pokemon detail: %s %s", id, err.Error())
		return
	}

	// Get Pokemon Sprite
	pokeID := pokeDetail.ID
	pokeSprite := h.pokemonLogic.GetSpriteURLByID(pokeID)

	// Map Pokemon Abilities
	pokeAbilities := []string{}
	for _, a := range pokeDetail.Abilities {
		pokeAbilities = append(pokeAbilities, a.Ability.Name)
	}

	// Map Pokemon Types
	pokeTypes := []string{}
	for _, t := range pokeDetail.Types {
		pokeTypes = append(pokeTypes, t.Type.Name)
	}

	// Map Pokemon Stats
	pokeStats := []models.PokemonStats{}
	for _, s := range pokeDetail.Stats {
		pokeStat := models.PokemonStats{
			Name:     s.Stat.Name,
			BaseStat: s.BaseStat,
			Effort:   s.Effort,
		}
		pokeStats = append(pokeStats, pokeStat)
	}

	pokemonDetail := models.PokemonDetail{
		ID:        h.pokemonLogic.FormatID(pokeID),
		Name:      pokeDetail.Name,
		Sprite:    pokeSprite,
		Height:    pokeDetail.Height,
		Weight:    pokeDetail.Weight,
		Abilities: pokeAbilities,
		Types:     pokeTypes,
		Stats:     pokeStats,
	}

	// Send data to template
	data := make(map[string]interface{})
	data["pokemon_detail"] = pokemonDetail
	h.render.RenderTemplate(w, "page.detail.gohtml", &models.TemplateData{
		Title: id,
		Data:  data,
	})
}

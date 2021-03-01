package handlers

import (
	"html/template"
	"net/http"

	"github.com/wisnuanggoro/pokedex-web-go/models/pokemon"
)

type handler struct {
	templates  *template.Template
	pokemonSvc pokemon.Service
}

type HomeHandler interface {
	IndexList(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(templates *template.Template, pokemonSvc pokemon.Service) HomeHandler {
	return &handler{
		templates:  templates,
		pokemonSvc: pokemonSvc,
	}
}

func (h *handler) IndexList(w http.ResponseWriter, r *http.Request) {
	resp, err := h.pokemonSvc.GetPokemonCardList(0, 0)
	if err != nil {
		return
	}
	h.templates.ExecuteTemplate(w, "index.gohtml", resp)
}

package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wisnuanggoro/pokedex-web-go/models/pokemon"
)

type detailHandler struct {
	templates  *template.Template
	pokemonSvc pokemon.Service
}

type DetailHandler interface {
	DetailPage(w http.ResponseWriter, r *http.Request)
}

func NewDetailHandler(templates *template.Template, pokemonSvc pokemon.Service) DetailHandler {
	return &detailHandler{
		templates:  templates,
		pokemonSvc: pokemonSvc,
	}
}

func (h *detailHandler) DetailPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id := vars["id"]
	fmt.Println("masuk1")
	_, err := h.pokemonSvc.GetPokemonByID(vars["id"])
	if err != nil {
		log.Println(err)
	}
	fmt.Println("masuk2")
	h.templates.ExecuteTemplate(w, "detail.gohtml", nil)
}

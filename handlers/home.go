package handlers

import (
	"fmt"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/models"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type handler struct {
	render       render.Render
	errorHandler ErrorHandler
}

type HomeHandler interface {
	CardList(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(render render.Render, errorHandler ErrorHandler) HomeHandler {
	return &handler{
		render:       render,
		errorHandler: errorHandler,
	}
}

func (h *handler) CardList(w http.ResponseWriter, r *http.Request) {
	p, err := pokeapi.Resource("pokemon", 0, 10000)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	data := make(map[string]interface{})
	data["results"] = p.Results
	h.render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{
		Data: data,
	})
}

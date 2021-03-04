package handlers

import (
	"net/http"

	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type detailHandler struct {
	render       render.Render
	errorHandler ErrorHandler
}

type DetailHandler interface {
	DetailPage(w http.ResponseWriter, r *http.Request)
}

func NewDetailHandler(render render.Render, errorHandler ErrorHandler) DetailHandler {
	return &detailHandler{
		render:       render,
		errorHandler: errorHandler,
	}
}

func (h *detailHandler) DetailPage(w http.ResponseWriter, r *http.Request) {
	h.render.RenderTemplate(w, "page.detail.gohtml", nil)
}

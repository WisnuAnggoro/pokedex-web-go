package handlers

import (
	"net/http"

	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type detailHandler struct {
	render render.Render
}

type DetailHandler interface {
	DetailPage(w http.ResponseWriter, r *http.Request)
}

func NewDetailHandler(render render.Render) DetailHandler {
	return &detailHandler{
		render: render,
	}
}

func (h *detailHandler) DetailPage(w http.ResponseWriter, r *http.Request) {
	h.render.RenderTemplate(w, "detail.page.gohtml", nil)
}

package handlers

import (
	"net/http"

	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type errorHandler struct {
	render render.Render
}

type ErrorHandler interface {
	ShowErrorPage(w http.ResponseWriter, r *http.Request, status int)
}

func NewErrorHandler(render render.Render) ErrorHandler {
	return &errorHandler{
		render: render,
	}
}

func (h *errorHandler) ShowErrorPage(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		h.render.RenderTemplate(w, "notfound.page.gohtml", nil)
	}
}

package handler

import (
	"html/template"
	"net/http"
)

type handler struct {
	templates *template.Template
}

type HomeHandler interface {
	IndexList(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler(templates *template.Template) HomeHandler {
	return &handler{
		templates: templates,
	}
}

func (h *handler) IndexList(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.gohtml", nil)
}

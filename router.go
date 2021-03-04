package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handlers"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

func InitRouter(render render.Render, cfg *config.Config) *mux.Router {
	// Create template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	cfg.TemplateCache = templateCache

	// Initialize handlers
	errorHandler := handlers.NewErrorHandler(render)
	homeHandler := handlers.NewHomeHandler(*cfg, render, errorHandler)
	detailHandler := handlers.NewDetailHandler(render, errorHandler)

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.HomePage).Methods("GET")
	r.HandleFunc("/detail/{id}", detailHandler.DetailPage).Methods("GET")

	// Initialize static folder
	fsImage := http.FileServer(http.Dir("./views/assets/img"))
	r.PathPrefix("/assets/img/").Handler(http.StripPrefix("/assets/img/", fsImage))
	fsCSS := http.FileServer(http.Dir("./views/assets/css"))
	r.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", fsCSS))
	fsJs := http.FileServer(http.Dir("./views/assets/js"))
	r.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/js/", fsJs))

	return r
}

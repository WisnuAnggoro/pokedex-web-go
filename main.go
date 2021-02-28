package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handler"
	"github.com/wisnuanggoro/pokedex-web-go/pokemon"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Set cache expiration for pokeapi
	cacheDuration, _ := strconv.Atoi(cfg.CacheDuration)
	pokeapi.CacheSettings.CustomExpire = time.Duration(cacheDuration)

	// Initialiaze templates
	templates := template.Must(template.ParseGlob("templates/*.gohtml"))

	// Initialize Service
	pokemonSvc := pokemon.NewService(cfg.PokemonSprite)

	// Initialize handlers
	homeHandler := handler.NewHomeHandler(templates, pokemonSvc)
	detailHandler := handler.NewDetailHandler(templates, pokemonSvc)

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.IndexList).Methods("GET")
	r.HandleFunc("/detail/{id}", detailHandler.DetailPage).Methods("GET")

	// Initialize static folder
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Run server
	http.ListenAndServe(":"+cfg.Port, r)
}

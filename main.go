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
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Set cache expiration for pokeapi
	cacheDuration, _ := strconv.Atoi(cfg.CacheDuration)
	pokeapi.CacheSettings.CustomExpire = time.Duration(cacheDuration)

	// Initialiaze templates
	templates := template.Must(template.ParseGlob("templates/*.gohtml"))

	// Initialize handlers
	homeHandler := handler.NewHomeHandler(templates)

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.IndexList).Methods("GET")

	// Run server
	http.ListenAndServe(":"+cfg.Port, r)
}

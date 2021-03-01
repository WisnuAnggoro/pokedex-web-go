package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handlers"
	"github.com/wisnuanggoro/pokedex-web-go/models/pokemon"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Set cache expiration for pokeapi
	cacheDuration, _ := strconv.Atoi(cfg.CacheDuration)
	pokeapi.CacheSettings.CustomExpire = time.Duration(cacheDuration)

	// Initialize templates
	templates := template.Must(template.ParseGlob("views/*.gohtml"))

	// Initialize Service
	pokemonSvc := pokemon.NewService(cfg.PokemonSprite)

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler(templates, pokemonSvc)
	detailHandler := handlers.NewDetailHandler(templates, pokemonSvc)

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.IndexList).Methods("GET")
	r.HandleFunc("/detail/{id}", detailHandler.DetailPage).Methods("GET")

	// Initialize static folder
	fsImage := http.FileServer(http.Dir("./views/assets/img"))
	r.PathPrefix("/assets/img/").Handler(http.StripPrefix("/assets/img/", fsImage))
	fsCSS := http.FileServer(http.Dir("./views/assets/css"))
	r.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", fsCSS))
	fsJs := http.FileServer(http.Dir("./views/assets/js"))
	r.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/css/", fsJs))

	// Run server
	fmt.Println(fmt.Sprintf("Starting application on port %s", cfg.Port))
	http.ListenAndServe(":"+cfg.Port, r)
}

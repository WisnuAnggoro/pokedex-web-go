package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handlers"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Set cache expiration for PokeAPI
	cacheDuration, _ := strconv.Atoi(cfg.PokeAPICacheDuration)
	pokeapi.CacheSettings.CustomExpire = time.Duration(cacheDuration)

	// Initialize utilities
	renderUtil := render.NewRender(&cfg)

	// Create template cache
	templateCache, err := renderUtil.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	cfg.TemplateCache = templateCache
	fmt.Println(fmt.Printf("%v", cfg.TemplateCache))

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler(renderUtil)
	detailHandler := handlers.NewDetailHandler(renderUtil)

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.CardList).Methods("GET")
	r.HandleFunc("/detail/{id}", detailHandler.DetailPage).Methods("GET")

	// Initialize static folder
	fsImage := http.FileServer(http.Dir("./views/assets/img"))
	r.PathPrefix("/assets/img/").Handler(http.StripPrefix("/assets/img/", fsImage))
	fsCSS := http.FileServer(http.Dir("./views/assets/css"))
	r.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", fsCSS))
	fsJs := http.FileServer(http.Dir("./views/assets/js"))
	r.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/js/", fsJs))

	// Run server
	fmt.Println(fmt.Sprintf("Starting application on port %s", cfg.Port))
	http.ListenAndServe(":"+cfg.Port, r)
}

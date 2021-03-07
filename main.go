package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handlers"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
	"github.com/wisnuanggoro/pokedex-web-go/utils/router"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Split PokemonSprite
	cfg.PokemonSprites = strings.Split(cfg.PokemonSprite, ",")

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

	// Initialize handlers
	errorHandler := handlers.NewErrorHandler(renderUtil)
	homeHandler := handlers.NewHomeHandler(cfg, renderUtil, errorHandler)
	detailHandler := handlers.NewDetailHandler(renderUtil, errorHandler)

	// Initialize Router
	routerUtil := router.NewRouter(errorHandler, homeHandler, detailHandler)
	router := routerUtil.InitRouter(renderUtil, &cfg)

	// Run server
	fmt.Println(fmt.Sprintf("Starting application on port %s", cfg.Port))
	http.ListenAndServe(":"+cfg.Port, router)
}

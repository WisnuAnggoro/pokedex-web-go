package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/wisnuanggoro/pokedex-web-go/config"
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

	// Initialize Router
	router := InitRouter(renderUtil, &cfg)

	// Run server
	fmt.Println(fmt.Sprintf("Starting application on port %s", cfg.Port))
	http.ListenAndServe(":"+cfg.Port, router)
}

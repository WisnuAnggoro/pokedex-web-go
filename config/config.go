package config

import (
	"html/template"

	"github.com/kelseyhightower/envconfig"
)

// Config struct to implement model of this service's configuration
type Config struct {
	// Setting port for gin
	Port string `envconfig:"PORT" default:"8080"`

	// PokeAPI Setting
	PokeAPICacheDuration string `envconfig:"POKEAPI_CACHE_DURATION" default:"3600"`

	// Pokemon Sprite
	PokemonSprite  string `envconfig:"POKEMON_SPRITE" default:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/%v.svg,https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%v.png"`
	PokemonSprites []string

	// Rendering HTML Template
	UseTemplateCache bool `envconfig:"USE_TEMPLATE_CACHE" default:"false"`
	TemplateCache    map[string]*template.Template
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct to implement model of this service's configuration
type Config struct {
	// Setting port for gin
	Port string `envconfig:"PORT" default:"8080"`

	// Cache
	CacheDuration string `envconfig:"CACHE_DURATION" default:"0"`

	// Pokemon Sprite
	PokemonSprite string `envconfig:"POKEMON_SPRITE" default:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%v.png"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

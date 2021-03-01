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
	PokeAPICacheDuration string `envconfig:"POKEAPI_CACHE_DURATION" default:"0"`

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

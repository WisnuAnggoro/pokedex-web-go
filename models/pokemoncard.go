package models

type PokemonCard struct {
	ID          int      `json:"id"`
	IDFormatted string   `json:"id_formatted"`
	Name        string   `json:"name"`
	Sprite      string   `json:"sprite"`
	Types       []string `json:"types"`
}

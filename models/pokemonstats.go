package models

type PokemonStats struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
}

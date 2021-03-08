package models

type PokemonDetail struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Sprite    string         `json:"sprite"`
	Height    int            `json:"height"`
	Weight    int            `json:"weight"`
	Abilities []string       `json:"abilities"`
	Types     []string       `json:"types"`
	Stats     []PokemonStats `json:"stats"`
}

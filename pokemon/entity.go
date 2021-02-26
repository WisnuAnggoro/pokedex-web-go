package pokemon

type PokemonCard struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Sprite    string   `json:"sprite"`
	SpriteAlt string   `json:"sprite_alt"`
	Types     []string `json:"types"`
}

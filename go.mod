module github.com/vinwong7/pokedexCLI

go 1.23.4

require github.com/vinwong7/pokedexCLI/internal/pokeapi v0.0.0
replace github.com/vinwong7/pokedexCLI/internal/pokeapi => ./internal/pokeapi
require github.com/vinwong7/pokedexCLI/internal/pokecache v0.0.0
replace github.com/vinwong7/pokedexCLI/internal/pokecache => ./internal/pokecache
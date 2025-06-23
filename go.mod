module Pokedex

go 1.24.2

require internal/pokeApi v0.0.0
replace internal/pokeApi => ./internal/pokeApi
require internal/pokecache v0.0.0
replace internal/pokecache => ./internal/pokecache
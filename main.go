package main

import (
	"fmt"

	"github.com/xSaCh/dweep/pkg/api"
)

const (
	OMDB_API_KEY = "5d0b46f6"
)

func main() {
	a := api.NewOmdbApi(OMDB_API_KEY)
	res := a.SearchFilmByTitle("Godzilla", nil)
	fmt.Printf("res: %v\n", res[0])
}

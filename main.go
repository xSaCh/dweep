package main

import (
	"fmt"
	"reflect"

	"github.com/xSaCh/dweep/pkg/api"
	"github.com/xSaCh/dweep/pkg/mocks"
)

const (
	OMDB_API_KEY = "5d0b46f6"
	TMDB_API_KEY = "a3ca43df787ec6b692b7e1e2d53a65ec"
)

func main() {
	ta := api.NewTmdbApi(TMDB_API_KEY)
	m := *ta.GetMovie("550")
	// for _, v := range m.Keywords {
	// 	fmt.Println(v)
	// }
	// fmt.Println(m)
	// fmt.Println(mocks.MovieFilms[0])

	CompareStructField(m, mocks.MovieFilms[0])
}

func CompareStructField(s1, s2 interface{}) {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	if v1.Type() != v2.Type() {
		panic("Cannot compare structs of different types")
	}

	for i := 0; i < v1.NumField(); i++ {
		field1 := v1.Field(i).Interface()
		field2 := v2.Field(i).Interface()

		if !reflect.DeepEqual(field1, field2) {
			fmt.Printf("%v != %v\n", field1, field2)
		}
	}

}

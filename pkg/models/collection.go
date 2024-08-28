package models

import "time"

// Collection
//     collection_id
//     owner_id
//     is_public

//     films
//         film_id
//         added_on
//         added_by

type FilmCollection struct {
	CollectionId int                  `json:"collection_id"`
	OwnerId      int                  `json:"owner_id"`
	IsPublic     bool                 `json:"is_public"`
	Films        []FilmCollectionItem `json:"films"`
}

type FilmCollectionItem struct {
	FilmId  int       `json:"film_id"`
	AddedOn time.Time `json:"added_on"`
	AddedBy int       `json:"added_by"`
}

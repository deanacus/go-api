package album

import (
	"github.com/deanacus/go-api/repositories"
)

type Album struct {
	repositories.Base
	Name     string `json:"name"`
	ArtistID int    `json:"artist_id"`
}

var albumsCollection = repositories.NewCollection[Album]()

func GetOne(id int) (Album, error) {
	return albumsCollection.GetOne(id)
}

func GetMany() []Album {
	return albumsCollection.GetMany()
}

func Create(data []byte) (Album, error) {
	return albumsCollection.Create(data)
}

func Update(id int, data []byte) (Album, error) {
	return albumsCollection.Update(id, data)
}

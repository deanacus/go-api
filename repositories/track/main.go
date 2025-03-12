package track

import "github.com/deanacus/go-api/repositories"

type Track struct {
	repositories.Base
	Name     string `json:"name"`
	ArtistID int    `json:"artist_id"`
	AlbumID  int    `json:"album_id"`
}

var tracksCollection = repositories.NewCollection[Track]()

func GetOne(id int) (Track, error) {
	return tracksCollection.GetOne(id)
}

func GetMany() []Track {
	return tracksCollection.GetMany()
}

func Create(data []byte) (Track, error) {
	return tracksCollection.Create(data)
}

func Update(id int, data []byte) (Track, error) {
	return tracksCollection.Update(id, data)
}

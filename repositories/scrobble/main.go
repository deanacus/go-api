package scrobble

import "github.com/deanacus/go-api/repositories"

type Scrobble struct {
	repositories.Base
	ArtistID   int    `json:"artist_id"`
	AlbumID    int    `json:"album_id"`
	TrackID    int    `json:"track_id"`
	ListenedAt string `json:"listened_at"`
}

var tracksCollection = repositories.NewCollection[Scrobble]()

func GetOne(id int) (Scrobble, error) {
	return tracksCollection.GetOne(id)
}

func GetMany() []Scrobble {
	return tracksCollection.GetMany()
}

func Create(data []byte) (Scrobble, error) {
	return tracksCollection.Create(data)
}

func Update(id int, data []byte) (Scrobble, error) {
	return tracksCollection.Update(id, data)
}

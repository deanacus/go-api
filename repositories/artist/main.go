package artist

import "github.com/deanacus/go-api/repositories"

type Artist struct {
	repositories.Base
	Name string `json:"name"`
	Type string `json:"type"`
}

var artistsCollection = repositories.NewCollection[Artist]()

func GetOne(id int) (Artist, error) {
	return artistsCollection.GetOne(id)
}

func GetMany() []Artist {
	return artistsCollection.GetMany()
}

func Create(data []byte) (Artist, error) {
	return artistsCollection.Create(data)
}

func Update(id int, data []byte) (Artist, error) {
	return artistsCollection.Update(id, data)
}

package oldartist

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var artists = make(map[int]Artist)
var currentID = 0
var artistsMu sync.Mutex

func GetOne(id int) (Artist, error) {
	artistsMu.Lock()
	defer artistsMu.Unlock()
	artistsMu.Lock()
	defer artistsMu.Unlock()
	artist, ok := artists[id]
	if !ok {
		return artist, errors.New("not found")
	}
	return artist, nil
}

func GetMany() []Artist {
	artistsMu.Lock()
	defer artistsMu.Unlock()

	as := make([]Artist, 0, len(artists))
	for _, a := range artists {
		as = append(as, a)
	}
	return as
}

func Create(data []byte) (Artist, error) {
	var artist Artist
	artistsMu.Lock()
	defer artistsMu.Unlock()

	err := json.Unmarshal(data, &artist)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return artist, err
	}

	currentID++
	artist.ID = currentID
	artists[artist.ID] = artist

	return artist, nil
}

func Update(id int, data []byte) (Artist, error) {
	artistsMu.Lock()
	defer artistsMu.Unlock()

	artist, ok := artists[id]
	if !ok {
		log.Println("Article not found")
		return artist, errors.New("not found")
	}

	err := json.Unmarshal(data, &artist)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return artist, errors.New("invalid data")
	}

	artist.ID = id
	artists[id] = artist
	return artist, nil
}

package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/deanacus/go-api/repositories/artist"
)

func handleArtistGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	path, ok := strings.CutPrefix(req.URL.Path, "/artists/")

	// Get many
	if !ok {
		artists := artist.GetMany()
		json.NewEncoder(res).Encode(artists)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("Invalid artist ID error: %v", err)
		http.Error(res, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	body, err := artist.GetOne(id)
	if err != nil {
		http.Error(res, "artist not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(body)
}

func handleArtistPost(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Request reading error: %v", err)
		http.Error(res, "Unable to read request body", http.StatusBadRequest)
		return
	}

	artist, err := artist.Create(body)
	if err != nil {
		http.Error(res, "Error creating artist", http.StatusBadRequest)
		return
	}

	log.Printf("Artist created: %v", artist)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(artist)

}

func ArtistsHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		handleArtistGet(res, req)
		return
	}

	if req.Method == http.MethodPost {
		handleArtistPost(res, req)
		return
	}

	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

func ArtistHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		handleArtistGet(res, req)
		return
	}

	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

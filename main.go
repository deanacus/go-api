package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/deanacus/go-api/routes"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.SetFlags(0)
		prefix := fmt.Sprintf("%s %s %s: ", time.Now().Format(time.RFC3339)[0:19], r.Method, r.URL.Path)
		log.SetPrefix(prefix)
		f(w, r)
	}
}

func main() {
	http.HandleFunc("/articles", logging(routes.ArticlesHandler))
	http.HandleFunc("/articles/", logging(routes.ArticleHandler))
	http.HandleFunc("/artists", logging(routes.ArtistsHandler))
	http.HandleFunc("/artists/", logging(routes.ArtistHandler))
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

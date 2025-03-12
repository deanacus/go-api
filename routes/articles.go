package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/deanacus/go-api/repositories/article"
)

func handleArticleGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	path, ok := strings.CutPrefix(req.URL.Path, "/articles/")

	// Get many
	if !ok {
		articles := article.GetMany()
		json.NewEncoder(res).Encode(articles)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("Invalid article ID error: %v", err)
		http.Error(res, "Invalid article ID", http.StatusBadRequest)
		return
	}

	body, err := article.GetOne(id)
	if err != nil {
		http.Error(res, "Article not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(body)
}

func handleArticlePost(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Request reading error: %v", err)
		http.Error(res, "Unable to read request body", http.StatusBadRequest)
		return
	}

	article, err := article.Create(body)
	if err != nil {
		log.Printf("Article creation error: %v", err)
		http.Error(res, "Error creating article", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(article)
}

func ArticlesHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		handleArticleGet(res, req)
		return
	}

	if req.Method == http.MethodPost {
		handleArticlePost(res, req)
		return
	}

	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

func ArticleHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		handleArticleGet(res, req)
		return
	}

	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

package article

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	articles   = make(map[int]Article)
	currentID  = 0
	articlesMu sync.Mutex
)

func GetOne(id int) (Article, error) {
	articlesMu.Lock()
	defer articlesMu.Unlock()
	article, ok := articles[id]
	if !ok {
		return article, errors.New("not found")
	}
	return article, nil
}

func GetMany() []Article {
	articlesMu.Lock()
	defer articlesMu.Unlock()

	as := make([]Article, 0, len(articles))
	for _, p := range articles {
		as = append(as, p)
	}
	return as
}

func Create(data []byte) (Article, error) {
	var article Article
	articlesMu.Lock()
	defer articlesMu.Unlock()

	err := json.Unmarshal(data, &article)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return article, err
	}

	currentID++
	article.ID = currentID
	articles[article.ID] = article

	return article, nil
}

func Update(id int, data []byte) (Article, error) {
	articlesMu.Lock()
	defer articlesMu.Unlock()

	article, ok := articles[id]
	if !ok {
		log.Println("Article not found")
		return article, errors.New("not found")
	}

	err := json.Unmarshal(data, &article)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
		return article, errors.New("invalid data")
	}

	articles[id] = article
	return article, nil
}

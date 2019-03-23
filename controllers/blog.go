package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"play-with-go/entities"
	"strconv"

	"github.com/gorilla/mux"
)

// init blog var as slice Blog stuct
var articles []entities.Article

func Mock() {
	// mock data
	articles = append(articles, entities.Article{ID: "1", Title: "Just test article", Body: "This is just test article", Category: &entities.Category{ID: "1", Title: "uncategory"}})
	articles = append(articles, entities.Article{ID: "2", Title: "Just test article 2", Body: "This is just test article 2", Category: &entities.Category{ID: "1", Title: "uncategory"}})
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&entities.Article{})
}

func StoreArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article entities.Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(10000))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var article entities.Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = strconv.Itoa(rand.Intn(10000))
			articles = append(articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

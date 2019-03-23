package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"play-with-go/entities"

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
	workDirectory, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}
	template, _ := template.ParseFiles(workDirectory + "/views/article.html")

	w.WriteHeader(http.StatusOK)
	template.Execute(w, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workDirectory, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}
	template, _ := template.ParseFiles(workDirectory + "/views/article-detail.html")
	for _, item := range articles {
		if item.ID == params["id"] {
			article := map[string]interface{}{"Item": item}
			w.WriteHeader(http.StatusOK)
			template.Execute(w, article)
		}
	}
	template.Execute(w, &entities.Article{})
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Article create view
}

func StoreArticle(w http.ResponseWriter, r *http.Request) {
	// Article create action
}

func EditArticle(w http.ResponseWriter, r *http.Request) {
	// Article update view
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// Article update action
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Delete action
}

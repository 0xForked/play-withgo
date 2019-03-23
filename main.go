package main

import (
	"net/http"

	"play-with-go/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// initialize the route
	rotuer := mux.NewRouter()

	// call mock
	controllers.Mock()

	// route handler (endpoints)
	rotuer.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)
	rotuer.HandleFunc("/", controllers.Home).Methods("GET")
	rotuer.HandleFunc("/blogs", controllers.GetArticles).Methods("GET")
	rotuer.HandleFunc("/blogs/{id}", controllers.GetArticle).Methods("GET")
	rotuer.HandleFunc("/blogs", controllers.StoreArticle).Methods("POST")
	rotuer.HandleFunc("/blogs/{id}", controllers.UpdateArticle).Methods("PUT")
	rotuer.HandleFunc("/blogs/{id}", controllers.DeleteArticle).Methods("DELETE")

	http.ListenAndServe(":8080", rotuer)
}

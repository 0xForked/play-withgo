package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to aasumtro</h1>")
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my blog</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404</h1>")
}

func main() {
	route := mux.NewRouter()
	route.NotFoundHandler = http.HandlerFunc(notFound)
	route.HandleFunc("/", home)
	route.HandleFunc("/blog", blog)
	http.ListenAndServe(":8080", route)
}

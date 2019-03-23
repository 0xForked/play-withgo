package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>Welcome Home</h1>")
	fmt.Fprint(w, "<a href='http://localhost:8080/api/v1/blogs'>Blog</a>")
}

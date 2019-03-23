package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Message": "Guest!"}

	workDirectory, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}
	template, _ := template.ParseFiles(workDirectory + "/views/home.html")

	w.WriteHeader(http.StatusOK)
	template.Execute(w, data)
}

package controllers

import (
	"go-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(w, "index", products)
}

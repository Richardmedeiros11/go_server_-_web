package controllers

import (
	"import/models"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

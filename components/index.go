package components

import (
	"api-go/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.RecuperaProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

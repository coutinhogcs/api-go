package routes

import (
	"api-go/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(":8000", nil)
}

package controllers

import (
	"api-go/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.RecuperaProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}
		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		idstr := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		quantidadeStr := r.FormValue("quantidade")

		id, err := strconv.Atoi(idstr)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}
		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}
		models.AtualizaProduto(id, nome, descricao, preco, quantidade)

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/henrikkorsgaard/bff-ftoda/ftoda"
)

func GetLovforslag(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/lovforslag.html")
	if err != nil {
		panic(err)
	}
	apiService := ftoda.NewFTODAService()
	sager, err := apiService.GetLovforslag(0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "list", sager)
}

func GetLovforslagById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("./templates/lovforslag.html")
	if err != nil {
		panic(err)
	}
	apiService := ftoda.NewFTODAService()
	//101403
	sag, err := apiService.GetLovforslagById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "lovforslag", sag)
}

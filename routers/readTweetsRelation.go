package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FernandoTorresBautista/gocourse-project/db"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página cómo entero mayor a 0", http.StatusBadRequest)
		return
	}
	response, ok := db.ReadTweetsFollowers(IDUsuario, page)
	if !ok {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// QueryRelation ...
func QueryRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.ResponseQueryRelation

	status, err := db.QueryRelation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

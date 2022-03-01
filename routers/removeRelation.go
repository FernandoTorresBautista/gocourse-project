package routers

import (
	"net/http"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// RemoveRelation ...
func RemoveRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relación", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

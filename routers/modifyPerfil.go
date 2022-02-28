package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// ModifyRegister ...
func ModifyRegister(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModifyRegister(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevmente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

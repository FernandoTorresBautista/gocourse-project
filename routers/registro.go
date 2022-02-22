package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// Register is the function to create a record of user on db
func Register(w http.ResponseWriter, r *http.Request) {

	var t = models.Usuario{}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Error, el email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Error, el password de usuario requiere al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := db.CheckIfUserExist(t.Email)
	if encontrado {
		http.Error(w, "Error, ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error intentando realizar el registro de usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error insertano el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FernandoTorresBautista/gocourse-project/db"
	"github.com/FernandoTorresBautista/gocourse-project/models"
)

// SaveTweet ...
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Ocurrio un error al leer los datos del usuario "+err.Error(), 400)
		return
	}

	register := models.SaveTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

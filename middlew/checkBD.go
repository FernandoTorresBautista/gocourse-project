package middlew

import (
	"net/http"

	"github.com/FernandoTorresBautista/gocourse-project/db"
)

// CheckBD middleware to get the status of DB
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

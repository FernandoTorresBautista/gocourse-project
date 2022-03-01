package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FernandoTorresBautista/gocourse-project/middlew"
	"github.com/FernandoTorresBautista/gocourse-project/routers"
)

// Handlers ...
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidateJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidateJWT(routers.ModifyRegister))).Methods("PUT")

	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.CheckBD(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckBD(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.AddRelation))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.RemoveRelation))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.QueryRelation))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

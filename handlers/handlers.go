package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Jhairofc/gotwitter/middleware"
	"github.com/Jhairofc/gotwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers conec  */
func Handlers() {
	router := mux.NewRouter()
	//Endpoints
	//Registro de Usuarios
	router.HandleFunc("/registro", middleware.Checkdb(routers.RegisterUs)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

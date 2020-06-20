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
	//Login de Usuarios
	router.HandleFunc("/login", middleware.Checkdb(routers.Login)).Methods("POST")
	//Procesar el Token del Login
	router.HandleFunc("/perfil", middleware.Checkdb(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	//Actualizar datos del usuario
	router.HandleFunc("/actualizar", middleware.Checkdb(middleware.ValidateJWT(routers.UpdateUs))).Methods("PUT")
	//Registrar un Tweet
	router.HandleFunc("/tweet", middleware.Checkdb(middleware.ValidateJWT(routers.RegisterTweet))).Methods("POST")
	//Leer mis Tweets
	router.HandleFunc("/mytweets", middleware.Checkdb(middleware.ValidateJWT(routers.MyTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

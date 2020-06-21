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
	//Borrar un Tweet
	router.HandleFunc("/borrarTweet", middleware.Checkdb(middleware.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")
	//Subir Avatar
	router.HandleFunc("/subirAvatar", middleware.Checkdb(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	//Obtener Avatar
	router.HandleFunc("/obtenerAvatar", middleware.Checkdb(routers.GetAvatar)).Methods("GET")
	//Subir Banner
	router.HandleFunc("/subirBanner", middleware.Checkdb(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")
	//Obtener Banner
	router.HandleFunc("/obtenerBanner", middleware.Checkdb(routers.GetBanner)).Methods("GET")
	//Guardar una relacion seguidor-siguiendo
	router.HandleFunc("/seguir", middleware.Checkdb(middleware.ValidateJWT(routers.Follow))).Methods("POST")
	//dejar de seguir a un usuario
	router.HandleFunc("/unfollow", middleware.Checkdb(middleware.ValidateJWT(routers.Unfollow))).Methods("DELETE")
	//Buscar a una persona para saber si le esta siguiendo o no
	router.HandleFunc("/follower", middleware.Checkdb(middleware.ValidateJWT(routers.Follower))).Methods("GET")
	//Lista usuarios que sigo y los que no sigo
	router.HandleFunc("/users", middleware.Checkdb(middleware.ValidateJWT(routers.AllUsers))).Methods("GET")
	//Lista de tweets de los usuarios que sigue el usuario logueado
	router.HandleFunc("/tweets", middleware.Checkdb(middleware.ValidateJWT(routers.Tweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

package routers

import (
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
)

//DeleteTweet metodo ruta para capturar ids y direccionar a borrar el tweet de un usuario
func DeleteTweet(res http.ResponseWriter, req *http.Request) {
	idTweet := req.URL.Query().Get("id")
	if len(idTweet) == 0 {
		http.Error(res, "El parametro id es obligatorio", 400)
		return
	}
	error := db.DeleteMyTweet(ID, idTweet)
	if error != nil {
		http.Error(res, "Error, no se ha podido borrar el tweet "+error.Error(), 400)
		return
	}
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
}

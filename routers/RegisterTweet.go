package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//RegisterTweet metodo ruta para capturar los datos del cliente y guardar en db
func RegisterTweet(res http.ResponseWriter, req *http.Request) {
	var tweet models.Tweet
	error := json.NewDecoder(req.Body).Decode(&tweet)
	if error != nil {
		http.Error(res, "No se ha podido obtener los datos del tweet"+error.Error(), 400)
		return
	}
	//Agregar userID y fecha al tweet y guardar el tweet en db
	tweet.UserID = ID
	tweet.Fecha = time.Now()
	_, status, err := db.SaveTweet(tweet)
	if err != nil {
		http.Error(res, "Error, no se ha podido guardar el tweet"+error.Error(), 400)
		return
	}
	//Si el status es false
	if status == false {
		http.Error(res, "No se ha podido guardar el tweet, intente nuevamente", 400)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

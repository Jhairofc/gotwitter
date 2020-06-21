package routers

import (
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//Unfollow metodo ruta para dejar de seguir a otro usuario
func Unfollow(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(res, "El campo id es necesario", 400)
		return
	}
	var unfollow models.Follow
	unfollow.UserID = ID
	unfollow.FollowingID = id
	status, error := db.DeleteFollow(unfollow)
	if error != nil || status == false {
		http.Error(res, "Ha ocurrido un error al hacer unfollow "+error.Error(), 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

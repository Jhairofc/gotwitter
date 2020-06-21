package routers

import (
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//Follow ruta para guardar una relacion seguidor-siguiendo
func Follow(res http.ResponseWriter, req *http.Request) {
	followingID := req.URL.Query().Get("id")
	if len(followingID) == 0 {
		http.Error(res, "El parametro ID es obligatorio", 400)
		return
	}
	var follow models.Follow
	follow.UserID = ID
	follow.FollowingID = followingID
	status, error := db.SaveFollow(follow)
	if error != nil || status == false {
		http.Error(res, "Ha ocurrido un error al guardar la relacion "+error.Error(), 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}

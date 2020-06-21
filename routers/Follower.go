package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//Follower metodo ruta para saber si una persona sigue a otra
func Follower(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(res, "El parametro ID es necesario", 400)
		return
	}
	var statusFollower models.Follower
	var follow models.Follow
	follow.UserID = ID
	follow.FollowingID = id
	status, error := db.GetFollower(follow)
	if error != nil || status == false {
		statusFollower.Status = false
	} else {
		statusFollower.Status = true
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(statusFollower)
}

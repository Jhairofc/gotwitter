package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
)

//ViewProfile ruta para extraer el ID del usuario y verificar su perfil
func ViewProfile(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(res, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	profile, error := db.SearchProfile(id)
	if error != nil {
		http.Error(res, "ha ocurrido un error al buscar al usuario"+error.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(profile)
}

package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jhairofc/gotwitter/db"
)

//AllUsers metodo ruta para obtener todos los usuarios o los seguidos por el usuario logueado
func AllUsers(res http.ResponseWriter, req *http.Request) {
	pageURL := req.URL.Query().Get("page")
	search := req.URL.Query().Get("search")
	_type := req.URL.Query().Get("type")
	if len(pageURL) == 0 || len(_type) == 0 {
		http.Error(res, "No todos los parametros son correctos", 400)
		return
	}
	_page, error := strconv.Atoi(pageURL)
	if error != nil {
		http.Error(res, "Error al leer el parámetro pagina "+error.Error(), 400)
		return
	}
	page := int64(_page)
	results, status := db.ReadUsers(ID, page, search, _type)
	if status == false {
		http.Error(res, "Ha ocurrido un error al lista los usuarios", 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

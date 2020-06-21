package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jhairofc/gotwitter/db"
)

//Tweets metodo ruta para obtener los tweets de los seguidores de una persona
func Tweets(res http.ResponseWriter, req *http.Request) {
	if len(req.URL.Query().Get("page")) == 0 {
		http.Error(res, "El parametro pagina es obligatorio", 400)
		return
	}
	_page, error := strconv.Atoi(req.URL.Query().Get("page"))
	if error != nil {
		http.Error(res, "El parametro pagina no tiene el formato correcto "+error.Error(), 400)
		return
	}
	page := int64(_page)
	results, status := db.ReadTweets(ID, page)
	if status == false {
		http.Error(res, "ha ocurrido un error al listar los tweets", 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

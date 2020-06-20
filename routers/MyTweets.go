package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jhairofc/gotwitter/db"
)

//MyTweets ruta para obtener los tweets de un usuario
func MyTweets(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(res, "El parametro ID es obligatorio", 400)
		return
	}
	if len(req.URL.Query().Get("pagina")) == 0 {
		http.Error(res, "El parametro pagina es obligatorio", 400)
		return
	}
	//Transformar el parametro URL de string a int con Atoi
	_page, error := strconv.Atoi(req.URL.Query().Get("pagina"))
	if error != nil {
		http.Error(res, "Error, el parametro pagina no tiene el formato correcto"+error.Error(), 400)
		return
	}
	//Pasar de int a int64
	page := int64(_page)
	results, status, err := db.ReadMyTweets(id, page)
	if err != nil {
		http.Error(res, "Error, no se ha podido leer los tweets: "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(res, "Error, no se ha podido leer los tweets", 400)
		return
	}
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

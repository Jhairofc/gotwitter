package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Jhairofc/gotwitter/db"
)

//GetBanner ruta para obtener el banner de un usuario
func GetBanner(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(res, "El parámetro ID es obligatorio", 400)
		return
	}
	profile, error := db.SearchProfile(id)
	if error != nil {
		http.Error(res, "Error, usuario no encontrado "+error.Error(), 400)
		return
	}
	//Abrir el archivo
	osFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(res, "Ha ocurrido un error al abrir el archivo "+error.Error(), 400)
		return
	}
	//Copiar el archivo en el res
	_, err = io.Copy(res, osFile)
	if err != nil {
		http.Error(res, "Ha ocurrido un error al copiar el archivo "+error.Error(), 400)
		return
	}

}

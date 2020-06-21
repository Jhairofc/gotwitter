package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//UploadBanner ruta para subir el archivo del banner del usuario
func UploadBanner(res http.ResponseWriter, req *http.Request) {
	file, handler, error := req.FormFile("banner")
	if error != nil {
		http.Error(res, "ha ocurrido un error al recibir el banner "+error.Error(), 400)
		return
	}
	extension := strings.Split(handler.Filename, ".")[1]
	path := "uploads/banners/" + ID + "." + extension
	//Subir a memoria el archivo en la carpeta path
	osFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(res, "Error al subir el archivo del banner "+error.Error(), 400)
		return
	}
	//Grabar el archivo en la carpeta path
	_, err = io.Copy(osFile, file)
	if err != nil {
		http.Error(res, "Error al copiar el archivo del banner "+error.Error(), 400)
		return
	}
	//Grabar el nombre del banner en db
	var user models.User
	var status bool
	user.Banner = ID + "." + extension
	status, err = db.UpdateUser(user, ID)
	if err != nil || status == false {
		http.Error(res, "Error al grabar el banner en la db "+error.Error(), 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}

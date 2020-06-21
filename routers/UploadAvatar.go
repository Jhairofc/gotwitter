package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//UploadAvatar ruta para subir el archivo del avatar del usuario
func UploadAvatar(res http.ResponseWriter, req *http.Request) {
	file, handler, error := req.FormFile("avatar")
	if error != nil {
		http.Error(res, "ha ocurrido un error al recibir el avatar "+error.Error(), 400)
		return
	}
	extension := strings.Split(handler.Filename, ".")[1]
	path := "uploads/avatars/" + ID + "." + extension
	//Subir a memoria el archivo en la carpeta path
	osFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(res, "Error al subir el archivo del avatar "+error.Error(), 400)
		return
	}
	//Grabar el archivo en la carpeta path
	_, err = io.Copy(osFile, file)
	if err != nil {
		http.Error(res, "Error al copiar el archivo del avatar "+error.Error(), 400)
		return
	}
	//Grabar el nombre del avatar en db
	var user models.User
	var status bool
	user.Avatar = ID + "." + extension
	status, err = db.UpdateUser(user, ID)
	if err != nil || status == false {
		http.Error(res, "Error al grabar el avatar en la db "+error.Error(), 400)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}

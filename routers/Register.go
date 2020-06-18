package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//RegisterUs Ruta para guardar el usuario
func RegisterUs(res http.ResponseWriter, req *http.Request) {
	var user models.User
	error := json.NewDecoder(req.Body).Decode(&user)
	//Valida si los datos enviados son erroneos
	if error != nil {
		http.Error(res, "Error al recibir los datos del usuario"+error.Error(), 400)
		return
	}
	//Valida si el correo es valido
	if len(user.Email) == 0 {
		http.Error(res, "Error, el email no puede ser nulo", 400)
		return
	}
	//Valida si la contrasena es valida
	if len(user.Password) < 6 {
		http.Error(res, "Error, la contrasena debe tener más de 6 caracteres", 400)
		return
	}
	//Valida si el usuario ya existe
	_, existe, _ := db.ExisteUser(user.Email)
	if existe == true {
		http.Error(res, "Error, Usuario ya existe", 400)
		return
	}
	//Guardar el usuario en db
	_, status, err := db.SaveUs(user)
	if err != nil {
		http.Error(res, "Error, Hubo un problema al registrar el usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(res, "Error, Usuario no ha podido ser registrado", 400)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

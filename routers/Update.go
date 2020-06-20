package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
)

//UpdateUs metodo ruta para obtener los datos usuario que llega del front
func UpdateUs(res http.ResponseWriter, req *http.Request) {
	var user models.User
	error := json.NewDecoder(req.Body).Decode(&user)
	if error != nil {
		http.Error(res, "Error, no se ha podido capturar los datos del usuario"+error.Error(), 400)
		return
	}
	//Enviar los datos a la db
	status, err := db.UpdateUser(user, ID)
	if err != nil {
		http.Error(res, "Error, no se ha podido actualizar los datos del usuario"+error.Error(), 400)
		return
	}
	//Si no dio error, pero no actualizo la informacion
	if status == false {
		http.Error(res, "Error, cero registros afectados", 400)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

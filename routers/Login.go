package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/jwt"
	"github.com/Jhairofc/gotwitter/models"
)

//Login Ruta para realizar el login del usuario
func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	var user models.User
	error := json.NewDecoder(req.Body).Decode(&user)
	//Validaciones
	if error != nil {
		http.Error(res, "Usuario i/o Password inválidos"+error.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(res, "El email es requerido", 400)
		return
	}
	model, exist := db.LoginIntent(user.Email, user.Password)
	if exist == false {
		http.Error(res, "Usuario i/o Password inválidos", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(model)
	if err != nil {
		http.Error(res, "Ha ocurrido un error al generar el token"+err.Error(), 400)
		return
	}
	result := models.LoginResponse{
		Token: jwtKey,
	}
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(result)

	//Creacion de una Cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(res, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

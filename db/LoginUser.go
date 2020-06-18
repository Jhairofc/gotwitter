package db

import (
	"github.com/Jhairofc/gotwitter/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginIntent realizamos la verificacion del login del usuario
func LoginIntent(email string, password string) (models.User, bool) {
	//Utilizando la funcion ExisteUser
	user, error, _ := ExisteUser(email)
	if error == false {
		return user, false
	}
	//Verificamos la password
	passDB := []byte(user.Password)
	passLogin := []byte(password)
	err := bcrypt.CompareHashAndPassword(passDB, passLogin)
	if err != nil {
		return user, false
	}
	return user, true
}

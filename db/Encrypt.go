package db

import "golang.org/x/crypto/bcrypt"

//EncryptPassword funcion utilizando bcrypt para encriptar la password
func EncryptPassword(pass string) (string, error) {
	//Esfuerzo es 2ˆeffort, son las veces que la clave sera encriptada se recomienda 2ˆ8 = 256 veces
	effort := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), effort)
	return string(bytes), err
}

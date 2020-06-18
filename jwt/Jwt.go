package jwt

import (
	"time"

	"github.com/Jhairofc/gotwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generamos el token utilizando jwt
func GenerateJWT(user models.User) (string, error) {
	privateKey := []byte("EstaEsMiClavePrivadaDePrueba")
	//Un JWT se forma de 3 partes: header, data o payload y Firma.

	//Pasamos toda la data que estara disponible en el jwt en formato json
	payload := jwt.MapClaims{
		"nombre":    user.Nombre,
		"apellido":  user.Apellido,
		"fechanac":  user.FechaNac,
		"email":     user.Email,
		"avatar":    user.Avatar,
		"banner":    user.Banner,
		"biografia": user.Biografia,
		"ubicacion": user.Ubicacion,
		"sitioweb":  user.Sitioweb,
		"id":        user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	//Agregamos el Header y el payload al token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//Firmamos al token utilizando la clave privada
	tokenSig, err := token.SignedString(privateKey)
	if err != nil {
		return tokenSig, err
	}
	return tokenSig, nil
}

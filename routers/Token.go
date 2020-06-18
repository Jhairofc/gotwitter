package routers

import (
	"errors"
	"strings"

	"github.com/Jhairofc/gotwitter/db"
	"github.com/Jhairofc/gotwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email Variable publica Exportable
var Email string

//ID Variable publica Exportable
var ID string

//ProcessToken Ruta para extraer los datos del token
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	privateKey := []byte("EstaEsMiClavePrivadaDePrueba")
	claim := &models.Claim{}

	//Verificamos si el token que llego es valido, se utiliza la palabra Bearer como separador
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, string(""), errors.New("Formato del token invalido")
	}
	//Limpiamos de espacion en blanco al token
	token = strings.TrimSpace(splitToken[1])
	//Mediantes una sintaxis porpia de jwt extraemos el payload del token a el modelo Claim
	tkn, err := jwt.ParseWithClaims(token, claim, func(_token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	//Si no hay errores verificamos que el usuario del token exista
	if err == nil {
		_, existe, _ := db.ExisteUser(claim.Email)
		if existe == true {
			//En Variables publicas guardamos tanto el mail como el ID, para que puedan ser ocupadas en todo el proyecto
			Email = claim.Email
			ID = claim.ID.Hex()
		}
		return claim, existe, ID, nil
	}
	//Si hubo un error 'err', verificamos si se produjo un fallo en obtener los datos con el ParseWithClaims
	if !tkn.Valid {
		return claim, false, string(""), errors.New("Token inválido")
	}
	//Si el token fue valido pero aun asi produjo el error, devolvemos el error
	return claim, false, string(""), err
}

package db

import (
	"context"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveUs Metodo para gusrdar el usuario en db
func SaveUs(user models.User) (string, bool, error) {
	//Se define un subContexto para esta rutina, el cual se ejecutara en un tiempo no mayor a 15seg
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//Despues de terminar la ejecucion se limpia el conexto con la funcion cancel()
	defer cancel()

	//Conexion con la db y la coleccion del usuario
	db := Mongodb.Database("gotwitter")
	col := db.Collection("users")

	//Encriptar la contrasena del usuario
	user.Password, _ = EncryptPassword(user.Password)
	//Guardar el usuario en db
	result, err := col.InsertOne(ctx, user)

	//Validaciones
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

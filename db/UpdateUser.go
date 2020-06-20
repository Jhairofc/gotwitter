package db

import (
	"context"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateUser metodo que guarda en db al usuario actualizado
func UpdateUser(user models.User, id string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("users")
	//Pasar el id del usuario de string a ObjectID(asi esta guardado en db)
	objID, _ := primitive.ObjectIDFromHex(id)
	//Como es un update se tiene que armar el where
	where := bson.M{"_id": bson.M{"$eq": objID}} //$eq = equals
	//Para mandar a actualizar los datos del usuario se comprueba que campos viene con datos ya que desde el Front solo se envia los campos que han sufrido modificaciones
	model := make(map[string]interface{}) //Creamos un slice clave valor de tipo Interface que contendra los campos a modificar
	if len(user.Nombre) > 0 {
		model["nombre"] = user.Nombre
	}
	if len(user.Apellido) > 0 {
		model["apellido"] = user.Apellido
	}
	model["fechanac"] = user.FechaNac
	if len(user.Avatar) > 0 {
		model["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		model["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		model["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		model["ubicacion"] = user.Ubicacion
	}
	if len(user.Sitioweb) > 0 {
		model["sitioweb"] = user.Sitioweb
	}
	//Para enviar el modelo a actualizar en Mongo se envia con la siguiente sintaxis
	modelMongo := bson.M{
		"$set": model,
	}
	_, err := col.UpdateOne(ctx, where, modelMongo)
	if err != nil {
		return false, err
	}
	return true, nil
}

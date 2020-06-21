package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadUsers metodo para obtener todos los usuarios de twitter o en base a una busqueda
func ReadUsers(id string, page int64, search string, _type string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("users")
	//Si se hace una busqueda de usuarios por nombres
	condition := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	//Where o limitantes que se pondra a la consulta
	_options := options.Find()
	_options.SetSkip((page - 1) * 20)
	_options.SetLimit(20)
	_options.SetProjection(
		bson.M{
			"email":     0,
			"password":  0,
			"banner":    0,
			"biografia": 0,
			"ubicacion": 0,
			"sitioWeb":  0,
		},
	)
	var results []*models.User
	//Consultar a la db
	cursor, error := col.Find(ctx, condition, _options)
	if error != nil {
		fmt.Println(error.Error())
		return results, false
	}
	var found, include bool
	//Recorremos el cursor para mostrar todos los usuarios excepto el usuario que consulta
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(error.Error())
			return results, false
		}
		//Verificar si el usuario es seguido por el usuario que realiza la consulta
		var follow models.Follow
		follow.UserID = id
		follow.FollowingID = user.ID.Hex()

		include = false
		found, err = GetFollower(follow)
		//Si es nuevo usuario y no es seguido por el usuarioLogin se incluye
		if found == false && _type == "new" {
			include = true
		}
		//Si es un usuario que yo sigo tbn incluir
		if found == true && _type == "follow" {
			include = true
		}
		//Si esque por algun motivo el usuario se sigue a si mismo, no listar
		if follow.FollowingID == id {
			include = false
		}
		//Una vez validada anteriormente los usuarios se almacena en el slice
		if include == true {
			results = append(results, &user)
		}
	}
	//Verificamos si no hubo un error interno del cursor
	err := cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true

}

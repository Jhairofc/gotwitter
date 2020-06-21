package db

import (
	"context"
	"log"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadTweets metodo para obtener los tweets de las personas quew se sigue.
func ReadTweets(id string, page int64) ([]models.Tweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("relations")
	skip := (page - 1) * 20
	//Para trabajar con una consulta de multiples tablas se utiliza el framwerok Agregate el cual acepta condiciones
	//Where sera un slice el cual se agregue varios filtros antes de ejecutar la consulta
	where := make([]bson.M, 0) //Crea un slice tipo bson.M vacio
	//$match = es un operador mongodb que filtra las consultas en base al ID del usuario logueado
	where = append(where, bson.M{"$match": bson.M{"userID": id}})
	//$lookup = operador que realiza la accion JOIN con otra tabla de interes
	where = append(where, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",      //tabla a la que se une
			"localField":   "followingID", //Campo de la tabla raiz(relacion) por lo cual se va a unir a la tabla tweets
			"foreignField": "userID",      //Campo por el cual relacionar en la tabla Tweets
			"as":           "tweets",      //Un alias a la tabla tweets, le puse el mismo nombre
		}})
	//$unwind = con ejemplo: Sin $unwind {nombre:'jairo', pc:['macbook', 'hp']};
	//Con $unwind {nombre:'jairo', pc: macbook} {nombre:'jairo', pc: hp}
	where = append(where, bson.M{"$unwind": "$tweets"})
	//$sort = operador para ordenar la consulta 1 Asc -1 Desc
	where = append(where, bson.M{"$sort": bson.M{"tweets.fecha": -1}})
	//$skip = operador para limitar la paginacion de la consulta, pag1:20reg, pag2:40, pag3:60 etc...
	where = append(where, bson.M{"$skip": skip})
	//%limit = cantidad de registros que se debe consulta
	where = append(where, bson.M{"$limit": 20})
	//una vez que se tenga todas las condiciones de la consulta se utiliza la funcion agregate
	var results []models.Tweets
	cursor, error := col.Aggregate(ctx, where)
	if error != nil {
		log.Printf("Error:" + error.Error())
		return results, false
	}
	//Se obtiene todos los datos del cursor y se los pasa al modelo tweets
	error = cursor.All(ctx, &results)
	if error != nil {
		log.Printf("Error:" + error.Error())
		return results, false
	}
	return results, true

}

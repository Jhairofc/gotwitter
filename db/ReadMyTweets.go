package db

import (
	"context"
	"log"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadMyTweets metodo para obtener desde db los tweets de un usuario
func ReadMyTweets(id string, page int64) ([]*models.MyTweets, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("tweets")
	var results []*models.MyTweets

	condition := bson.M{
		"userID": id,
	}
	//Nose puede traer todos los tweets del usuario de golpe por lo que se realizar el proceso de la siguiente manera
	//Se trae muestra de 20 en 20
	//De acuerdo a la pagina del usuario Ej: pag0:20, pag1:40, pag2:60 etc....
	_options := options.Find()
	_options.SetLimit(20)
	_options.SetSort(bson.D{{Key: "fecha", Value: -1}}) //-1 significa ordenar descendentemente
	_options.SetSkip((page - 1) * 20)
	//Obtener datos de la base de acuerdo a las condiciones anteriormente definidad
	cursor, error := col.Find(ctx, condition, _options)
	if error != nil {
		log.Fatal("Ha ocurrido un error: " + error.Error())
		return results, false, error
	}
	//Si todo se obtuvo bien se llena los tweets en el slice results y ese se retorna en la funcion
	for cursor.Next(context.TODO()) {
		var tweet models.MyTweets
		err := cursor.Decode(&tweet)
		if err != nil {
			return results, false, err
		}
		//Llenamos cada tweet en el slice de tweets
		results = append(results, &tweet)
	}
	return results, true, nil
}

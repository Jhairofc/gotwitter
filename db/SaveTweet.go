package db

import (
	"context"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveTweet Metodo que guarda un tweet en la db
func SaveTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("tweets")
	//Formatemos el modelo del tweet para poder pasarlo a Mongodb
	model := bson.M{
		"userID":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}
	result, error := col.InsertOne(ctx, model)
	if error != nil {
		return string(""), false, error
	}
	//Recuperamos el id del Tweet
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil

}

package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteMyTweet metodo para borrar un tweet en la db
func DeleteMyTweet(userID string, tweetID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("tweets")
	//Convert strint go ObjectID
	objID, _ := primitive.ObjectIDFromHex(tweetID)
	condition := bson.M{
		"_id":    objID,
		"userID": userID,
	}
	//Borrar de la base
	_, error := col.DeleteOne(ctx, condition)
	return error
}

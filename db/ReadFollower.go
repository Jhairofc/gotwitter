package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetFollower consultar usuarios a los que sigue un usuario
func GetFollower(follow models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("relations")
	//Condicion
	condition := bson.M{
		"userID":      follow.UserID,
		"followingID": follow.FollowingID,
	}
	var result models.Follow
	error := col.FindOne(ctx, condition).Decode(&result)
	if error != nil {
		fmt.Println("Error: ", error.Error())
		return false, error
	}
	return true, nil
}

package db

import (
	"context"
	"time"

	"github.com/Jhairofc/gotwitter/models"
)

//SaveFollow metodo para guardar
func SaveFollow(relation models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("relations")

	_, error := col.InsertOne(ctx, relation)
	if error != nil {
		return false, error
	}
	return true, nil
}

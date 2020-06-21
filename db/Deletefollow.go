package db

import (
	"context"
	"time"

	"github.com/Jhairofc/gotwitter/models"
)

//DeleteFollow dejar de seguir a otro usuario
func DeleteFollow(follow models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("relations")
	//Borrar de la base la relacion seguidor-siguiendo
	_, error := col.DeleteOne(ctx, follow)
	if error != nil {
		return false, error
	}
	return true, nil
}

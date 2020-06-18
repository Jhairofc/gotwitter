package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jhairofc/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchProfile Buscar el perfil del usuario en db
func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := Mongodb.Database("gotwitter")
	col := db.Collection("users")
	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condition).Decode(&profile)
	//Nunca se debe devolver desde el backend la password por lo que se setea en nil
	profile.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return profile, err
	}
	return profile, nil
}

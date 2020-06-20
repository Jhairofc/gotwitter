package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MyTweets modelo para obtener el listado de los tweets de un usuario
type MyTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID  string             `bson:"userID" json:"userID,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}

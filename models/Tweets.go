package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tweets modelo join para unir los tweets de todos los usuarios a los que el usuario logueado sigue
type Tweets struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID      string             `bson:"userID" json:"userID,omitempty"`
	FollowingID string             `bson:"followingID" json:"followingID,omitempty"`
	Tweets      struct {
		ID      string    `bson:"_id" json:"id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}

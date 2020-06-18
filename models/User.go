package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User modelo para la administracion de usuarios
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido  string             `bson:"apellido" json:"apellido,omitempty"`
	FechaNac  time.Time          `bson:"fechanac" json:"fechanac,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografia string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	Sitioweb  string             `bson:"sitioweb" json:"sitioweb,omitempty"`
}

package models

import "time"

//Tweet Modelo para manejo de tweets
type Tweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}

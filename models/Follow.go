package models

//Follow modelo para le relacion seguidor-siguiendo
type Follow struct {
	UserID      string `bson:"userID" json:"userID"`
	FollowingID string `bson:"followingID" json:"followingID"`
}

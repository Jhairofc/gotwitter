package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongodb tiene la conexion con la db
var Mongodb = mongoConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@cluster0-ojqwy.mongodb.net/test?retryWrites=true&w=majority")

//mongoConnect() permite realizar la conexion con MongoDB
func mongoConnect() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOptions)
	//Comprobamos si la conexion fue exitosa
	if error != nil {
		log.Fatal(error)
		return client
	}
	//Comprobamos si ademas de que la conexion fue exitosa, la db esta levantada y operativa
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
		return client
	}
	//Si todo fue OK retorna el client
	log.Println("Conexión exitosa")
	return client
}

//CheckConnectionDB Check Ping si la db esta lista
func CheckConnectionDB() int {
	error := Mongodb.Ping(context.TODO(), nil)
	if error != nil {
		return 0
	}
	return 1
}

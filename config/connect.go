package config

import (
	"context"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoUrl string = "mongodb+srv://csin45159:csin45159@cluster0.d0sf9oc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func Connect() *mongo.Client {
	mongoclient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatal("error while connecting to DB", err)
	}
	log.Printf("Connection successful")
	err = mongoclient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("error ping the db", err)
	}
	return mongoclient
}

func Connection() *MongoColletions {
	mongoclient := Connect()

	collection := mongoclient.Database("TASKMANAGMENT").Collection("tasks")
	return &MongoColletions{mongoColletion: collection}
}

package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name       string `json:"namee"`
	Userid     int    `json:"id"`
	Password   string `json:"password"`
	Technology string `json:"tech"`
	Edit       string `json:"edit"`
}
type Tip struct {
	Date       int    `json:"date"`
	Technology string `json:"technology"`
	Content    string `json:"content"`
}
type Addtip struct {
	Content    string `json:"content"`
	Technology string `json:"technology"`
}

func Createdatabase() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Tool := client.Database("Tools")

	collectionName := "GO"
	err = Tool.CreateCollection(ctx, collectionName)
	if err != nil {
		log.Fatal(err)
	}
	collectionName2 := "MongoDB"
	err = Tool.CreateCollection(ctx, collectionName2)
	if err != nil {
		log.Fatal(err)
	}
	collectionName3 := "Java"
	err = Tool.CreateCollection(ctx, collectionName3)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

}

package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	fmt.Println("Connected to MongoDB!")
	db = client.Database("test")

}

type Trainer struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	City string `bson:"city"`
}

type URLData struct {
	ShortURL  string `json:"short_url" bson:"short_url"`
	Exptime   int64  `json:"exp_time" bson:"exp_time"`
	OriginURL string `json:"origin_url" bson:"origin_url"`
}

func test() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("trainers")
	defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)

	ash := Trainer{"Ash", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	insert, err := collection.InsertOne(ctx, ash)

	fmt.Println("aaaa", insert.InsertedID)

}

// 存储url
func SaveURL(d *URLData) {
	collection := db.Collection("urls")
	insert, err := collection.InsertOne(context.TODO(), d)
	if err != nil {
		fmt.Println("insert failed")
	}
	fmt.Println(insert)
}

func GetURL(s string) {

	collection := db.Collection("urls")
	filter := bson.M{"short_url": s}
	if cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(1)); err != nil {
		fmt.Println("find fail")
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			fmt.Println("cursor fial")
		}
	}()

}

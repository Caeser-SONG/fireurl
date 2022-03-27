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
	if err != nil {
		fmt.Println(err)
		return
	}
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
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	insert, err := collection.InsertOne(ctx, d)
	if err != nil {
		fmt.Println("insert failed")
	}
	fmt.Println(insert)
}

func GetOriginURL(s string) (string, error) {
	collection := db.Collection("urls")
	filter := bson.M{"short_url": s}
	fmt.Printf("shortv = %s \n", s)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	sresult := collection.FindOne(ctx, filter)

	var result bson.M
	err := sresult.Decode(&result)
	if err != nil {
		fmt.Printf("no data in mongo with %s, err = %s", s, err)
		return "", fmt.Errorf("no data")
	}
	return result["origin_url"].(string), nil
}

func IsExist(s string) bool {
	collection := db.Collection("urls")
	filter := bson.M{"short_url": s}
	fmt.Printf("shortv = %s \n", s)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	sresult := collection.FindOne(ctx, filter)
	var result bson.M
	err := sresult.Decode(&result)
	if err != nil {
		return false
	}
	return true
}

package main

import (
	"context"
	"encoding/base64"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return client, cancel
}

// make async
func UpdateUserRFTokenHash(id string, newHash []byte) {
	coll := client.Database("mydb").Collection("users")

	res := base64.StdEncoding.EncodeToString(newHash)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "rf_token_hash", Value: res}}}}
	filter := bson.D{{Key: id}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func AddUser(id string, newHash []byte) {
	coll := client.Database("mydb").Collection("users")

	res := base64.StdEncoding.EncodeToString(newHash)
	_, err := coll.InsertOne(context.TODO(), User{ID: id, RFTokenHash: res})
	if err != nil {
		panic(err)
	}
}

func IsUserExist(id string) bool {
	coll := client.Database("mydb").Collection("users")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Err()

	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}

func IsHashExistsInDB(hash []byte) bool {
	coll := client.Database("mydb").Collection("users")

	res := base64.StdEncoding.EncodeToString(hash)
	err := coll.FindOne(context.TODO(), bson.D{{Key: "rf_token_hash", Value: res}}).Err()

	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}

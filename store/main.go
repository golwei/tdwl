package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://some-mongo:27017")) //client
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //ctx
	err = client.Connect(ctx)                                           //connect
	if err != nil {
		panic(err)
	}
	collection := client.Database("tdwlxx").Collection("goods")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	println(id)
	// ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	// cur, err := collection.Find(ctx, bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// do something with result....
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}

func C() {
	bson.D{"name", bson.D{"$in", bson.A{"li4", "zhangsan"}}}
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	trainers := []interface{}{misty, brock}

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://some-mongo:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	collection := client.Database("tdwlxx").Collection("goods")

	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	println(insertManyResult.InsertedIDs)

	//close
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func UpdateOne() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://some-mongo:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	collection := client.Database("tdwlxx").Collection("goods")

	filter := bson.D{{"name", "Ash"}}
	update := bson.D{{"$inc", bson.D{{"age", 1}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	//close
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
func FindOne() {
	var result Trainer
	filter := bson.D{{"name", "Ash"}}
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://some-mongo:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	// _ = client.Ping(context.TODO(), nil)
	collection := client.Database("tdwlxx").Collection("goods")

	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	}
	//close
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
func D() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://some-mongo:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	// _ = client.Ping(context.TODO(), nil)
	collection := client.Database("tdwlxx").Collection("goods")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	cur.Next()
}

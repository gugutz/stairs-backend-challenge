package dao

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DAO {
	Server string
	Database string
}

//Set the collection to use in mongo db
collection := client.Database("wines_db").Collection("wines")
host = "mongodb://localhost:27017"

// Set client options
client, err := mongo.NewClient(options.Client().ApplyURI({host}))
if err != nil { return err }

// Connect to MongoDB
func (dao *DAO) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	fmt.Println("Connected to MongoDB!")
	if err != nil { return err }
}

func (dao *DAO) GetAll (){
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}



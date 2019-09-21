package dao

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DAO struct {
	Server string
	Database string
	Collection string
}

// Connect to MongoDB
func (dao *DAO) Connect() {
	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI(dao.Server)
	if err != nil { return err }

	// connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	else {
		fmt.Println("Connected to MongoDB!")
	}
}

func (dao *DAO) GetAll (){
	//Set the collection to use in mongo db
	collection := client.Database(dao.Database).Collection(dao.Collection)
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



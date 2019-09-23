package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"github.com/gugutz/stairs-backend-challenge/models"
)

type DAO struct {
	Server string
	Database string
	Collection string
}

var client *mongo.Client
var collection *mongo.Collection

func (dao *DAO) CreateClient() (client *mongo.Client, error error) {
	dbclient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Error creating Mongo Client: %v \n", err)
	}
	client = dbclient
	return client, err
}

func (dao *DAO) GetContext(deadline time.Duration) (client context.Context, cancel context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), deadline*time.Second)
	if ctx != nil {
		return ctx, cancel
	} else {
		fmt.Printf("Error getting the context: %v \n", cancel)
	}
	return ctx, cancel
}

func (dao *DAO) Connect() error {
	client, err := dao.CreateClient()
	ctx, _ := dao.GetContext(10)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v \n", err)
		return err
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	collection = client.Database(dao.Database).Collection(dao.Collection)
	return err
}

func (dao *DAO) Create(wine models.Wine) (*mongo.InsertOneResult, error) {
	ctx, cancel := dao.GetContext(5)
	defer cancel()
	result, err := collection.InsertOne(ctx, wine)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func (dao *DAO) Get(id string) (models.Wine, error) {
	ctx, cancel := dao.GetContext(5)
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	var result models.Wine
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Printf("Error getting the document from the DB: %+v\n", err)
		log.Fatal(err)
	}
	return result, err
}
func (dao *DAO) GetAll() ([]*models.Wine, error) {
	var results []*models.Wine
	ctx, cancel := dao.GetContext(30)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error generating cursor: %v", err)
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var elem models.Wine
		err := cursor.Decode(&elem)
		if err != nil { log.Fatal(err) }
		results = append(results, &elem)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(ctx)
	return  results, err
}

func (dao *DAO) Update(id string, wine models.Wine) (*mongo.UpdateResult, error){
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	update := bson.D{
		{"$set", bson.D{
			{"name", wine.Name},
			{"harvest", wine.Harvest},
			{"country", wine.Country},
			{"ammount", wine.Ammount},
			{"description", wine.Description},
		}},
	}
	ctx, cancel := dao.GetContext(30)
	defer cancel()
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("Error updating document: ", err)
		log.Fatal(err)
	}
	return result, err
}

func (dao *DAO) DeleteOne(id string) (*mongo.DeleteResult, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Error deleting the document from the DB: %+v\n", err)
		log.Fatal(err)
	}
	return result, err
}


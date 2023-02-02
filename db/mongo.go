package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client *mongo.Client
}

// This function creates a new connection client with Mongo that will be used for all our db operations
func NewMongoStore() (*mongoStore, error) {
	uri := "mongodb://mongodb:27017/?maxPoolSize=20&w=majority"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	names, err := client.Database("spamhaus").ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return nil, err
	}

	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == "urls" {
			log.Printf("The urls collection exists!")
			return &mongoStore{Client: client}, nil
		}
	}

	log.Printf("The urls collection has been initialized")
	return &mongoStore{Client: client}, nil
}

func (m mongoStore) Get(id string) (map[string]interface{}, error) {

	coll := m.Client.Database("spamhaus").Collection("urls")
	filter := bson.M{"_id": id}
	var result map[string]interface{}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (m mongoStore) Update(id string, f map[string]interface{}) (map[string]interface{}, error) {

	col := m.Client.Database("spamhaus").Collection("urls")
	opts := options.Update().SetUpsert(true)
	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": id},
		bson.M{"$set": bson.M{"json": f}}, opts)

	if err != nil {
		return f, err
	}

	return f, nil
}

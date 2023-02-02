package db

import (
	"assignment/domain"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	MS mongoStore
}
type URLRequestInterface interface {
	GetBySize() ([]domain.URLEntry, error)
	GetByDate() ([]domain.URLEntry, error)
	GetMostSubmitted() ([]domain.URLEntry, error)
	Update(d domain.URLEntry) (*domain.URLEntry, error)
	Delete(d domain.URLEntry) error
}

func NewRepository(ms mongoStore) *Repository {
	return &Repository{
		MS: ms,
	}
}

func (r *Repository) GetBySize() ([]domain.URLEntry, error) {
	var err error
	var cur *mongo.Cursor
	var ctx = context.Background()
	entries := make([]domain.URLEntry, 0)
	col := r.MS.Client.Database("spamhaus").Collection("urls")

	aggregatePipeline := bson.A{}
	aggregatePipeline = append(aggregatePipeline, bson.D{{
		Key: "$addFields",
		Value: bson.D{{
			Key:   "size",
			Value: bson.D{{"$binarySize", bson.A{"$Data"}}},
		}},
	}})
	aggregatePipeline = append(aggregatePipeline, bson.D{{"$sort", bson.D{{"size", -1}}}})
	aggregatePipeline = append(aggregatePipeline, bson.D{{"$limit", 50}})

	opts := options.Aggregate()
	opts.SetBatchSize(50)
	if cur, err = col.Aggregate(ctx, aggregatePipeline, opts); err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		dom := domain.URLEntry{}
		if err := cur.Decode(&dom); err != nil {
			return nil, err
		}
		entries = append(entries, dom)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return entries, nil

}

func (r *Repository) GetByDate() ([]domain.URLEntry, error) {
	var err error
	var cur *mongo.Cursor
	var ctx = context.Background()
	entries := make([]domain.URLEntry, 0)
	col := r.MS.Client.Database("spamhaus").Collection("urls")

	opts := options.Find().SetSort(bson.M{"$natural": -1})
	opts.SetLimit(50)

	if cur, err = col.Find(ctx, bson.M{}, opts); err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		dom := domain.URLEntry{}
		if err := cur.Decode(&dom); err != nil {
			return nil, err
		}
		entries = append(entries, dom)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *Repository) GetMostSubmitted() ([]domain.URLEntry, error) {
	var err error
	var cur *mongo.Cursor
	var ctx = context.Background()
	entries := make([]domain.URLEntry, 0)
	col := r.MS.Client.Database("spamhaus").Collection("urls")

	aggregatePipeline := bson.A{}
	aggregatePipeline = append(aggregatePipeline, bson.D{{
		Key: "$addFields",
		Value: bson.D{{
			Key:   "count",
			Value: bson.D{{"$max", bson.A{"$SubmissionCount"}}},
		}},
	}})
	aggregatePipeline = append(aggregatePipeline, bson.D{{"$sort", bson.D{{"count", -1}}}})
	aggregatePipeline = append(aggregatePipeline, bson.D{{"$limit", 10}})
	opts := options.Aggregate()
	if cur, err = col.Aggregate(ctx, aggregatePipeline, opts); err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		dom := domain.URLEntry{}
		if err := cur.Decode(&dom); err != nil {
			return nil, err
		}
		entries = append(entries, dom)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *Repository) Update(d domain.URLEntry) (*domain.URLEntry, error) {
	col := r.MS.Client.Database("spamhaus").Collection("urls")
	// objectID, _ := primitive.ObjectIDFromHex(id)
	var ctx = context.Background()
	filter := bson.M{"WebsiteURL": d.WebsiteURL}
	var retrievedEntry domain.URLEntry

	err := col.FindOne(ctx, filter).Decode(&retrievedEntry)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	if retrievedEntry.WebsiteURL != "" {
		d.SubmissionCount = retrievedEntry.SubmissionCount
	}
	d.SubmissionCount++
	update := bson.D{
		{"$set", d},
	}

	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result := col.FindOneAndUpdate(ctx, filter, update, &opts)
	if result.Err() != nil {
		return nil, result.Err()
	}
	entry := domain.URLEntry{}

	err = result.Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil

}

func (r *Repository) Delete(d domain.URLEntry) error {
	var ctx = context.Background()
	col := r.MS.Client.Database("spamhaus").Collection("urls")
	filter := bson.M{"WebsiteURL": d.WebsiteURL}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Println("Delete result: ", res)
	return nil

}

/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type boxesRepository struct {
	BaseRepository
}

func NewBoxesRepository(db *mongo.Database) core.BoxRepositoryI {

	str, err := db.Collection("boxes").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"location": "2dsphere"},
		Options: options.Index().SetSphereVersion(3),
	})
	log.Println("INDEX", str, err)
	return &boxesRepository{
		BaseRepository{Col: db.Collection("boxes")},
	}
}

func (repo *boxesRepository) FindOneAndIncrement(result *core.Box, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return repo.Col.FindOneAndUpdate(context.Background(),
		bson.M{"_id": objId},
		bson.M{"$inc": bson.M{"opened": 1}}).
		Decode(result)
}

func (repo *boxesRepository) FindNearBoxes(result *[]core.Box, lat, lng float64) error {

	// Specify GeoJSON filter
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{lat, lng},
				},
				"$maxDistance": 5000, // max Distance in meters
			},
		},
	}

	cursor, err := repo.Col.Find(context.Background(), filter)
	if err != nil {
		return err
	}

	// Return all results & error
	return cursor.All(context.Background(), result)
}

// Returns Top-20 boxes
func (repo *boxesRepository) FindTopBoxes(result *[]core.Box) error {

	// Specify order by downloads filter with limit 20
	filter := bson.A{
		bson.M{"$sort": bson.M{
			"downloaded": -1,
			"opened":     -1},
		},
		bson.M{"$limit": 20},
	}

	cursor, err := repo.Col.Aggregate(context.Background(), filter)
	if err != nil {
		return err
	}

	// Return all results & error
	return cursor.All(context.Background(), result)
}

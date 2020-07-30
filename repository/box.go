/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type boxesRepository struct {
	BaseRepository
}

func NewBoxesRepository(db *mongo.Database) core.BoxRepositoryI {
	return &boxesRepository{
		BaseRepository{Col: db.Collection("boxes")},
	}
}

func (repo *boxesRepository) FindNearBoxes(result *[]core.Box) error {

	// Specify GeoJSON filter
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{-73.9667, 40.78},
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
		bson.M{"$sort": bson.M{"downloaded": -1}},
		bson.M{"$limit": 20},
	}

	cursor, err := repo.Col.Aggregate(context.Background(), filter)
	if err != nil {
		return err
	}

	// Return all results & error
	return cursor.All(context.Background(), result)
}

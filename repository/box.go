/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"github.com/MikaelLazarev/filebox-server/core"
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



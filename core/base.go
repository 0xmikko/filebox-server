/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	BaseModel struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		CreatedAt int64              `json:"createdAt"`
		UpdatedAt int64              `json:"updatedAt"`
	}

	BaseModelI interface {
		SetID(newId primitive.ObjectID)
		GetID() primitive.ObjectID
		SetCreatedAt()
		SetUpdatedAt()
	}

	BaseRepositoryI interface {
		FindOne(result interface{}, id string) error
		FindAll(result interface{}) error
		Create(item BaseModelI) error
		Save(item BaseModelI) error
	}
)

func (b *BaseModel) SetID(newId primitive.ObjectID) {
	b.ID = newId
}

func (b *BaseModel) GetID() primitive.ObjectID {
	return b.ID
}

func (b *BaseModel) SetCreatedAt() {
	b.CreatedAt = time.Now().Unix()
}

func (b *BaseModel) SetUpdatedAt() {
	b.CreatedAt = time.Now().Unix()
}

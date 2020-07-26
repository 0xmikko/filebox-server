/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	BaseModel struct {
		ID        string
		//CreatedAt time.Time  `json:"createdAt"`
		//UpdatedAt time.Time  `json:"updatedAt"`
	}

	BaseRepositoryI interface {
		FindOne(result interface{}, id string) error
		FindAll(result interface{}) error
		Create(item interface{}) error
		//Save(item interface{}) error
	}
)



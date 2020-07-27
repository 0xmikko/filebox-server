/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	Box struct {
		BaseModel
		IpfsID string
		Name   string
		Lat    float64
		Lng    float64
	}

	BoxRepositoryI interface {
		BaseRepositoryI
	}

	BoxServiceI interface {
		Create(tmpFilename, filename string) (*Box, error)
		FindBoxesAround() ([]Box, error)
		Retrieve(id string) (*Box, error)
	}
)

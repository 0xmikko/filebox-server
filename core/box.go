/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import p "github.com/MikaelLazarev/filebox-server/payload"

type (
	Box struct {
		BaseModel
		IPFSHash string
		Name     string
		Lat      float64
		Lng      float64
		Altitude float64
	}

	BoxRepositoryI interface {
		BaseRepositoryI
	}

	BoxServiceI interface {
		Create(request p.BoxCreateRequest, tmpFilename, filename string) (*Box, error)
		FindBoxesAround() ([]Box, error)
		Retrieve(id string) (*Box, error)
	}
)

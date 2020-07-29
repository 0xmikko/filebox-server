/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import p "github.com/MikaelLazarev/filebox-server/payload"

type (
	Box struct {
		BaseModel
		IPFSHash string  `json:"ipfs_hash"`
		Name     string  `json:"name"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
		Altitude float64 `json:"altitude"`
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

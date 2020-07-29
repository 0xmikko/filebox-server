/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import p "github.com/MikaelLazarev/filebox-server/payload"

type (
	Box struct {
		BaseModel
		IPFSHash   string  `json:"ipfsHash"`
		Name       string  `json:"name"`
		Lat        float64 `json:"lat"`
		Lng        float64 `json:"lng"`
		Altitude   float64 `json:"altitude"`
		Content    string  `json:"content"`
		Opened     int     `json:"opened"`
		Downloaded int     `json:"downloaded"`
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

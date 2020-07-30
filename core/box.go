/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	Location [2]float64

	Box struct {
		BaseModel
		IPFSHash   string   `json:"ipfsHash"`
		Name       string   `json:"name"`
		Location   Location `json:"location"`
		Altitude   float64  `json:"altitude"`
		Content    string   `json:"content"`
		Opened     int      `json:"opened"`
		Downloaded int      `json:"downloaded"`
	}

	BoxRepositoryI interface {
		BaseRepositoryI
		FindNearBoxes(lat, lng float64, result *[]Box) error
		FindTopBoxes(*[]Box) error
	}

	BoxServiceI interface {
		Create(req BoxCreateRequest, tmpFilename, filename string) (*Box, error)
		FindNearAndTopBoxes(req BoxListRequest) (*BoxListResponse, error)
		Retrieve(id string) (*Box, error)
	}
)

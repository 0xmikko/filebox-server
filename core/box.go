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
		FindNearBoxes(*[]Box) error
		FindTopBoxes(*[]Box) error
	}

	BoxServiceI interface {
		Create(request BoxCreateRequest, tmpFilename, filename string) (*Box, error)
		FindNearAndTopBoxes() (*BoxListResponse, error)
		Retrieve(id string) (*Box, error)
	}
)

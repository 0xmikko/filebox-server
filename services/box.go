/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
)

type boxService struct {
	repository core.BoxRepositoryI
}

func NewBoxService(repository core.BoxRepositoryI) core.BoxServiceI {
	return &boxService{repository: repository}
}

// Retrieves Box by ID
func (s *boxService) Retrieve(id string) (*core.Box, error) {
	var box core.Box
	if err := s.repository.FindOne(&box, id); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}
	return &box, nil
}

func (s *boxService) FindBoxesAround() ([]core.Box, error) {
	panic("implement me")
}

// Creates a new box and return it
func (s *boxService) Create() (*core.Box, error) {
	return nil, nil
}




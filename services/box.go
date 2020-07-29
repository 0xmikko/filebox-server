/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/MikaelLazarev/filebox-server/payload"
	"log"
	"os"
)

type boxService struct {
	repository core.BoxRepositoryI
	ipfs       core.IPFSRepositoryI
}

func NewBoxService(repository core.BoxRepositoryI, ir core.IPFSRepositoryI) core.BoxServiceI {
	return &boxService{
		repository: repository,
		ipfs:       ir}
}

// Retrieves Box by ID
func (s *boxService) Retrieve(id string) (*core.Box, error) {
	var box core.Box
	if err := s.repository.FindOne(&box, id); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}
	return &box, nil
}

// Find boxes around
func (s *boxService) FindBoxesAround() ([]core.Box, error) {
	panic("implement me")
}

// Creates a new box and return it
func (s *boxService) Create(boxDTO payload.BoxCreateRequest, tmpFilename, filename string) (*core.Box, error) {

	// Getting io.Reader by opening file
	r, err := os.Open(tmpFilename)
	if err != nil {
		return nil, err
	}

	newBox := core.Box{
		Name:     boxDTO.Name,
		Lat:      boxDTO.Lat,
		Lng:      boxDTO.Lng,
		Altitude: boxDTO.Altitude,
	}

	if err := s.repository.Create(&newBox); err != nil {
		return nil, err
	}

	log.Println(newBox)

	//go func() {
	ipfsHash, err := s.ipfs.AddFile(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(ipfsHash)
	newBox.IPFSHash = ipfsHash
	if err := s.repository.Save(&newBox); err != nil {
		log.Println(err)
	}
	//}()

	return &newBox, nil
}

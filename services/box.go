/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
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
	if err := s.repository.FindOneAndIncrement(&box, id); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}
	return &box, nil
}

func (s *boxService) Download(id string) (string, error) {
	var box core.Box
	if err := s.repository.FindOneAndIncrement(&box, id); err != nil {
		return "", err
	}

	return s.ipfs.GetFile(box.IPFSHash)
}

// Find near & top boxes around
func (s *boxService) FindNearAndTopBoxes(req core.BoxListRequest) (*core.BoxListResponse, error) {
	response := core.BoxListResponse{
		Near: make([]core.Box, 0, 0),
		Top:  make([]core.Box, 0, 0),
	}

	if err := s.repository.FindNearBoxes(&response.Near, req.Lat, req.Lng); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}
	if err := s.repository.FindTopBoxes(&response.Top); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}

	return &response, nil
}

// Creates a new box and return it
func (s *boxService) Create(boxDTO core.BoxCreateRequest, tmpFilename, filename string) (*core.Box, error) {

	// Getting io.Reader by opening file
	r, err := os.Open(tmpFilename)
	if err != nil {
		return nil, err
	}

	newBox := core.Box{
		Name:     boxDTO.Name,
		Location: [2]float64{boxDTO.Lat, boxDTO.Lng},
		Altitude: boxDTO.Altitude,
		Content:  "Owner doesn't provided description yet.",
	}

	if err := s.repository.Create(&newBox); err != nil {
		return nil, err
	}

	log.Println(newBox)

	go func() {
		ipfsHash, err := s.ipfs.AddFile(r)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(ipfsHash)
		newBox.IPFSHash = ipfsHash
		if err := s.repository.Save(&newBox); err != nil {
			log.Println(err)
		}
	}()

	return &newBox, nil
}

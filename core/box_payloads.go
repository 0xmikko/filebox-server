/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	// CREATE
	BoxCreateRequest struct {
		Name     string  `json:"name" validate:"required"`
		Lat      float64 `json:"lat" validate:"required"`
		Lng      float64 `json:"lng" validate:"required"`
		Altitude float64 `json:"altitude" validate:"required"`
	}

	// UPDATE
	BoxCreateResponse struct {
		ID string `json:"id"`
	}

	// LIST
	BoxListRequest struct {
		Lat float64 `form:"lat" validate:"required"`
		Lng float64 `form:"lng" validate:"required"`
	}

	BoxListResponse struct {
		Near []Box `json:"near"`
		Top  []Box `json:"top"`
	}
)

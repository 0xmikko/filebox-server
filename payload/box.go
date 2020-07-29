/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package payload

type (
	BoxCreateRequest struct {
		Name     string  `json:"name" validate:"required"`
		Lat      float64 `json:"lat" validate:"required"`
		Lng      float64 `json:"lng" validate:"required"`
		Altitude float64 `json:"altitude" validate:"required"`
	}

	// UPDATE Request
	BoxCreateResponse struct {
		ID string `json:"id"`
	}
)

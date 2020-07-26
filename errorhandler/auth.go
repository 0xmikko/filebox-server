/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

func ForbiddenError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Forbidden",
		Code:    http.StatusForbidden,
		Err:     e,
	}
}

func InvalidAuthorisationTokenError() ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Invalid authorization token",
		Code:    http.StatusForbidden,
		Err:     errors.New("invalid authorization token"),
	}
}

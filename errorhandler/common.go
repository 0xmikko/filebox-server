/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package errorhandler

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"runtime"
)

func HttpBadRequestError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Wrong request",
		Code:    http.StatusBadRequest,
		Err:     e,
	}
}

func HttpForbiddenRequestError() ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Forbidden",
		Code:    http.StatusBadRequest,
		Err:     errors.New("Forbidden"),
	}
}


func DBError(err error, msg string) error {
	if err == mongo.ErrNoDocuments {
		return ApiError{
			Module:  "",
			Message: msg,
			Code:    http.StatusNotFound,
			Err:     err,
		}
	}
	return UnknownError(err)
}

func UnknownError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Unknown error",
		Code:    http.StatusInternalServerError,
		Err:     e,
	}
}

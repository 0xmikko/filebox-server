/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package errorhandler

import (
	"fmt"
	"net/http"
	"runtime"
)

func UserNotError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "User not found",
		Code:    http.StatusNotFound,
		Err:     e,
	}
}


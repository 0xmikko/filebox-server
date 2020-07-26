/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package middlewares

import (
	"go.uber.org/fx"
)

var Module = fx.Invoke(InitJWTAuthMiddleware, InitHttpRedirect)

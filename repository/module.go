/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(ConnectDB,
	NewBoxesRepository,
	NewUsersRepository,
	NewIPFSClient)

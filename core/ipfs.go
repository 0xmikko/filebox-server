/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import "io"

type IPFSRepositoryI interface {
	AddFile(r io.Reader) (string, error)
	GetFile(ipfsHash string) ([]byte, error)
}
